// Package rain wraps the Open-Meteo Forecast and Archive APIs, which
// require no API key. Forecast (with past_days) gives the most recent
// observed rainfall the oracle needs daily; Archive gives the 3-year
// history the pricing engine needs.
package rain

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	defaultForecastURL = "https://api.open-meteo.com/v1/forecast"
	defaultArchiveURL  = "https://archive-api.open-meteo.com/v1/archive"
	requestTimeout     = 15 * time.Second
)

// DailyPoint is one day's rainfall, rounded to an integer millimeter per
// SCHEMA.md ("rainfall: integer millimeters, rounded half-up").
type DailyPoint struct {
	Date string // "YYYY-MM-DD", WIB calendar date
	Mm   int
}

// Client talks to Open-Meteo. ForecastURL and ArchiveURL are exported so
// tests can point them at an httptest server.
type Client struct {
	HTTPClient  *http.Client
	ForecastURL string
	ArchiveURL  string
}

func NewClient() *Client {
	return &Client{
		HTTPClient:  &http.Client{Timeout: requestTimeout},
		ForecastURL: defaultForecastURL,
		ArchiveURL:  defaultArchiveURL,
	}
}

type dailyResponse struct {
	Daily struct {
		Time             []string  `json:"time"`
		PrecipitationSum []float64 `json:"precipitation_sum"`
	} `json:"daily"`
}

func (c *Client) fetch(ctx context.Context, baseURL string, params url.Values) (dailyResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, baseURL+"?"+params.Encode(), nil)
	if err != nil {
		return dailyResponse{}, fmt.Errorf("rain.fetch: build request: %w", err)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return dailyResponse{}, fmt.Errorf("rain.fetch: request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return dailyResponse{}, fmt.Errorf("rain.fetch: unexpected status %d from %s", resp.StatusCode, baseURL)
	}

	var out dailyResponse
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return dailyResponse{}, fmt.Errorf("rain.fetch: decode: %w", err)
	}
	return out, nil
}

// ForecastDaily fetches the last two days plus today for one zone, in WIB
// calendar days. The Forecast API is used (rather than Archive) because
// Archive lags a few days behind real time.
func (c *Client) ForecastDaily(ctx context.Context, lat, lon float64) ([]DailyPoint, error) {
	params := url.Values{
		"latitude":      {formatCoord(lat)},
		"longitude":     {formatCoord(lon)},
		"daily":         {"precipitation_sum"},
		"timezone":      {"Asia/Jakarta"},
		"past_days":     {"2"},
		"forecast_days": {"1"},
	}
	resp, err := c.fetch(ctx, c.ForecastURL, params)
	if err != nil {
		return nil, fmt.Errorf("rain.ForecastDaily: %w", err)
	}
	return toDailyPoints(resp), nil
}

// ArchiveDaily fetches daily rainfall between fromDate and toDate
// ("YYYY-MM-DD", inclusive) from the Archive API.
func (c *Client) ArchiveDaily(ctx context.Context, lat, lon float64, fromDate, toDate string) ([]DailyPoint, error) {
	params := url.Values{
		"latitude":   {formatCoord(lat)},
		"longitude":  {formatCoord(lon)},
		"daily":      {"precipitation_sum"},
		"timezone":   {"Asia/Jakarta"},
		"start_date": {fromDate},
		"end_date":   {toDate},
	}
	resp, err := c.fetch(ctx, c.ArchiveURL, params)
	if err != nil {
		return nil, fmt.Errorf("rain.ArchiveDaily: %w", err)
	}
	return toDailyPoints(resp), nil
}

func formatCoord(v float64) string {
	return strconv.FormatFloat(v, 'f', -1, 64)
}

func toDailyPoints(resp dailyResponse) []DailyPoint {
	points := make([]DailyPoint, 0, len(resp.Daily.Time))
	for i, date := range resp.Daily.Time {
		var mm float64
		if i < len(resp.Daily.PrecipitationSum) {
			mm = resp.Daily.PrecipitationSum[i]
		}
		points = append(points, DailyPoint{Date: date, Mm: roundHalfUp(mm)})
	}
	return points
}

// roundHalfUp matches SCHEMA.md: "Rainfall: integer millimeters, rounded
// half-up." Negative values (not physically meaningful for precipitation)
// clamp to zero rather than propagating a sign error on-chain.
func roundHalfUp(v float64) int {
	if v <= 0 {
		return 0
	}
	return int(math.Floor(v + 0.5))
}

// PointForDate returns the point matching dateWIB ("YYYY-MM-DD"), or
// ok=false if that date isn't present in the response (e.g. the upstream
// API hasn't published it yet).
func PointForDate(points []DailyPoint, dateWIB string) (DailyPoint, bool) {
	for _, p := range points {
		if p.Date == dateWIB {
			return p, true
		}
	}
	return DailyPoint{}, false
}
