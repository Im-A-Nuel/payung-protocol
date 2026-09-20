package chain

import "time"

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
// computed arithmetically (no timezone database lookup needed): the UTC
// instant of that day's WIB midnight, formatted in UTC, is exactly that WIB
// calendar date.
func DateStringWIB(dayIndex uint32) string {
	unixAtWibMidnight := int64(dayIndex)*secondsPerDay - wibOffsetSeconds
	return time.Unix(unixAtWibMidnight, 0).UTC().Format("2006-01-02")
}
