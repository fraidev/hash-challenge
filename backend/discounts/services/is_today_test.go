package services

import (
	"testing"
	"time"
)

func TestIsToday(t *testing.T) {
	expectedDateCases := []time.Time{}

	currentDate := time.Now()
	expectedDateCases = append(expectedDateCases, currentDate)

	startCurrentDate := time.Date(currentDate.Year(), currentDate.Month(), currentDate.Day(), 0, 0, 0, 0, currentDate.Location())
	expectedDateCases = append(expectedDateCases, startCurrentDate)

	endCurrentDate := time.Date(currentDate.Year(), currentDate.Month(), currentDate.Day(), 23, 59, 0, 0, currentDate.Location())
	expectedDateCases = append(expectedDateCases, endCurrentDate)

	notExpectedDateCases := []time.Time{}

	notExpectedDateCases = append(notExpectedDateCases, time.Now().AddDate(0, 0, 1))
	notExpectedDateCases = append(notExpectedDateCases, time.Now().AddDate(0, 0, -1))

	for _, date := range expectedDateCases {
		if !IsToday(int(date.Month()), date.Day()) {
			t.Fatalf(date.Format("01-02"), "is not expected as today date")
		}
	}

	for _, date := range notExpectedDateCases {
		if IsToday(int(date.Month()), date.Day()) {
			t.Fatalf(date.Format("01-02"), "is expected as today date")
		}
	}
}
