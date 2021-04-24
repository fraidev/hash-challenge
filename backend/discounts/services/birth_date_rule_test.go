package services

import (
	"testing"
	"time"

	models "github.com/fraifelipe/hash-challenge/backend/discounts/models"
)

func TestBirthDateRuleValidate(t *testing.T) {
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
		user := models.User{DateOfBirth: date}
		rule := BirthDateRule{user: &user}
		if (!rule.Validate(&RuleEngine{})) {
			t.Fatalf(date.Format("01-02"), "is not expected as today date")
		}
	}

	for _, date := range notExpectedDateCases {
		user := models.User{DateOfBirth: date}
		rule := BirthDateRule{user: &user}
		if (rule.Validate(&RuleEngine{})) {
			t.Fatalf(date.Format("01-02"), "is expected as today date")
		}
	}
}

func TestChangePercentage(t *testing.T) {
	user := models.User{}
	rule := BirthDateRule{user: &user}
	if rule.ChangePercentage(3.0) != 8.0 {
		t.Fatalf("8.0 is expected Percentage")
	}

	if rule.ChangePercentage(2.0) != 7.0 {
		t.Fatalf("7.0 is expected Percentage")
	}

	if rule.ChangePercentage(1.0) != 6.0 {
		t.Fatalf("6.0 is expected Percentage")
	}
}
