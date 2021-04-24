package services

import (
	"testing"
)

func TestMaxPercentageRuleChangeValidate(t *testing.T) {

	var testCases = []struct {
		Input    float32
		Expected bool
	}{
		{0.0, false},
		{5.0, false},
		{10.0, false},
		{10.1, true},
		{11.0, true},
		{12.0, true},
	}

	rule := MaxPercentageRule{}

	for _, testeCase := range testCases {

		if rule.Validate(&RuleEngine{percentageDiscount: testeCase.Input}) != testeCase.Expected {
			t.Fatalf("%f input expected %t", testeCase.Input, testeCase.Expected)
		}
	}
}

func TestMaxPercentageRuleChangePercentage(t *testing.T) {

	var testCases = []struct {
		Input    float32
		Expected float32
	}{
		{0.0, 10.0},
		{5.0, 10.0},
		{10.0, 10.0},
		{10.1, 10.0},
		{11.0, 10.0},
		{12.0, 10.0},
	}

	rule := MaxPercentageRule{}

	for _, testeCase := range testCases {

		if rule.ChangePercentage(testeCase.Input) != testeCase.Expected {
			t.Fatalf("%f input expected %f", testeCase.Input, testeCase.Expected)
		}
	}
}
