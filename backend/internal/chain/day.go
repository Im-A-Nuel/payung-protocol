package chain

import (
	"fmt"
	"time"
)

// DayIndex and WeekIndex must compute identically to PayungPool.today() and
// week_index in Solidity. Never change this formula without updating both
// sides and their tests (see CLAUDE.md).

const (
	wibOffsetSeconds = int64(7 * 3600)
	secondsPerDay    = int64(86400)
)

// DayIndex returns floor((unixSeconds + 7h) / 86400), matching the
// contract's today() for a given unix timestamp.
func DayIndex(unixSeconds int64) uint32 {
	return uint32((unixSeconds + wibOffsetSeconds) / secondsPerDay)
}

// WeekIndex returns dayIndex / 7 (integer division), matching the
// contract's week_index used for the per-week payout cap.
func WeekIndex(dayIndex uint32) uint32 {
	return dayIndex / 7
}

// DateStringWIB returns the "YYYY-MM-DD" WIB calendar date for a day_index,
// computed arithmetically so no timezone database is needed.
//
// day_index d covers the unix range [d*86400-7h, (d+1)*86400-7h), and at any
// instant in it the WIB wall clock reads unix+7h. Rendering d*86400 in UTC
// therefore prints exactly the WIB calendar date that day_index names.
func DateStringWIB(dayIndex uint32) string {
	return time.Unix(int64(dayIndex)*secondsPerDay, 0).UTC().Format("2006-01-02")
}

// DayIndexFromDateWIB is the exact inverse of DateStringWIB. Callers that
// accept a date from the outside (the admin endpoints) must go through this
// rather than parsing by hand, so the two directions cannot drift apart.
func DayIndexFromDateWIB(tanggal string) (uint32, error) {
	t, err := time.Parse(time.DateOnly, tanggal)
	if err != nil {
		return 0, fmt.Errorf("chain.DayIndexFromDateWIB: %w", err)
	}
	return DayIndex(t.Unix()), nil
}
