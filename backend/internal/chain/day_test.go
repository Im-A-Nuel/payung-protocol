package chain

import (
	"testing"
	"time"
)

func TestDayIndex_MatchesReferenceFormula(t *testing.T) {
	cases := []int64{0, 1_800_000_000, 1_758_200_000, 4_102_444_800}
	for _, unix := range cases {
		got := DayIndex(unix)
		want := uint32((unix + 7*3600) / 86400)
		if got != want {
			t.Errorf("DayIndex(%d) = %d, want %d", unix, got, want)
		}
	}
}

// Three timestamps around WIB midnight (17:00 UTC the previous day, since
// WIB is UTC+7) must land on the correct day_index: just before, exactly
// at, and just after the boundary. Mirrors the Solidity fuzz test in
// contracts/test/PayungPool.t.sol so both sides agree on "which day".
func TestDayIndex_AroundWibMidnightBoundary(t *testing.T) {
	const anchor = int64(1_800_000_000)
	midnightWibUnix := anchor - (anchor % secondsPerDay) + secondsPerDay - wibOffsetSeconds

	dayBefore := DayIndex(midnightWibUnix - 1)
	dayAt := DayIndex(midnightWibUnix)
	dayAfter := DayIndex(midnightWibUnix + 1)

	if dayAt != dayBefore+1 {
		t.Errorf("day at boundary = %d, want dayBefore+1 = %d", dayAt, dayBefore+1)
	}
	if dayAfter != dayAt {
		t.Errorf("day after boundary = %d, want same as dayAt = %d", dayAfter, dayAt)
	}
}

// DateStringWIB and DayIndexFromDateWIB have to be exact inverses. They were
// not once: DateStringWIB subtracted the WIB offset from a day boundary and
// landed a day early, so the oracle looked up the wrong calendar day in the
// Open-Meteo response and settled a day nobody was covered for.
func TestDateStringWIB_RoundTripsWithDayIndexFromDateWIB(t *testing.T) {
	for _, day := range []uint32{0, 1, 19000, 20716, 20717, 20718, 25000} {
		tanggal := DateStringWIB(day)
		balik, err := DayIndexFromDateWIB(tanggal)
		if err != nil {
			t.Fatalf("DayIndexFromDateWIB(%q): %v", tanggal, err)
		}
		if balik != day {
			t.Errorf("day %d -> %q -> %d, want %d", day, tanggal, balik, day)
		}
	}
}

// An oracle run at 06:00 WIB on 17 Sep 2026 must process 16 Sep 2026, the day
// that just finished, and call it by that name.
func TestDateStringWIB_YesterdayFromSixAmWib(t *testing.T) {
	// 06:00 WIB on 2026-09-17 is 23:00 UTC on 2026-09-16.
	pagi := time.Date(2026, 9, 16, 23, 0, 0, 0, time.UTC).Unix()

	hariIni := DayIndex(pagi)
	if got := DateStringWIB(hariIni); got != "2026-09-17" {
		t.Errorf("hari ini = %q, want 2026-09-17", got)
	}
	if got := DateStringWIB(hariIni - 1); got != "2026-09-16" {
		t.Errorf("kemarin = %q, want 2026-09-16", got)
	}
}

func TestDayIndexFromDateWIB_RejectsGarbage(t *testing.T) {
	if _, err := DayIndexFromDateWIB("20 September 2026"); err == nil {
		t.Error("expected an error for a non ISO date")
	}
}

func TestWeekIndex_IsDayIndexDividedBySeven(t *testing.T) {
	cases := []uint32{0, 1, 6, 7, 8, 20832, 20839}
	for _, day := range cases {
		week := WeekIndex(day)
		if week != day/7 {
			t.Errorf("WeekIndex(%d) = %d, want %d", day, week, day/7)
		}
	}
}
