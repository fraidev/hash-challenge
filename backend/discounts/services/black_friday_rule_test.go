package services

import (
	"testing"
)

func TestBlackFridayRuleChangePercentage(t *testing.T) {

	var testCases = []struct {
		Input    float32
		Expected float32
	}{
		{0.0, 10.0},
		{5.0, 15.0},
		{10.0, 20.0},
	}

	rule := BlackFridayRule{}

	for _, testeCase := range testCases {

		if rule.ChangePercentage(testeCase.Input) != testeCase.Expected {
			t.Fatalf("%f input expected %f", testeCase.Input, testeCase.Expected)
		}
	}
}
