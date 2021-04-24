package services

import "time"

func IsToday(month int, day int) bool {
	currentDate := time.Now()
	return month == int(currentDate.Month()) && day == currentDate.Day()
}
