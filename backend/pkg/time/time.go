package time

import (
	"time"
)

const (
	Week           time.Duration = 24 * 7 * time.Hour
	YearsContract                = 1
	MonthsContract               = 0
	DaysContract                 = 0
)

func GetToday() time.Time {
	now := time.Now().UTC()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
}

func GetEndOfContract() time.Time {
	return time.Now().AddDate(YearsContract, MonthsContract, DaysContract)
}

func RelevantPeriod() time.Time {
	return GetToday().AddDate(0, -3, 0)
}

func CheckDateWeekLater(expectedDate time.Time) bool {
	return expectedDate.After(GetToday().Add(Week))
}

func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
