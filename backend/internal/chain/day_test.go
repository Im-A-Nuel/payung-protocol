package chain

import "testing"

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

func TestWeekIndex_IsDayIndexDividedBySeven(t *testing.T) {
	cases := []uint32{0, 1, 6, 7, 8, 20832, 20839}
	for _, day := range cases {
		week := WeekIndex(day)
		if week != day/7 {
			t.Errorf("WeekIndex(%d) = %d, want %d", day, week, day/7)
		}
	}
}
