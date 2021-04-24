package services

import (
	"testing"
)

func TestCalculateValueInCents(t *testing.T) {

	var testCases = []struct {
		PriceInCentsInput       uint64
		PercentageDiscountInput float32
		Expected                int32
	}{
		{50000, 10.0, 5000},
		{25000, 10.0, 2500},
		{10000, 5.0, 500},
		{1000, 1.0, 10},
	}

	for _, testcase := range testCases {
		if CalculateValueInCents(testcase.PriceInCentsInput, testcase.PercentageDiscountInput) != testcase.Expected {
			t.Fatalf("For price %d and percentage %f is expected %d value",
				testcase.PriceInCentsInput, testcase.PercentageDiscountInput, testcase.Expected)
		}
	}
}
