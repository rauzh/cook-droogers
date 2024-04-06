package time

import (
	"cookdroogers/internal/requests/sign_contract"
	"time"
)

const Week time.Duration = 24 * 7 * time.Hour

func GetToday() time.Time {
	now := time.Now().UTC()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
}

func GetEndOfContract() time.Time {
	return time.Now().AddDate(sign_contract.YearsContract, sign_contract.MonthsContract, sign_contract.DaysContract)
}

func RelevantPeriod() time.Time {
	return GetToday().AddDate(0, -3, 0)
}

func CheckDateWeekLater(expectedDate time.Time) bool {
	return expectedDate.After(GetToday().Add(Week))
}
