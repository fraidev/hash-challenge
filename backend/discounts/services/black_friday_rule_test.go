package services

import (
	"testing"
)

func TestBlackFridayRuleChangePercentage(t *testing.T) {
	rule := BlackFridayRule{}
	if rule.ChangePercentage(3.0) != 13.0 {
		t.Fatalf("13.0 is expected Percentage")
	}

	if rule.ChangePercentage(2.0) != 12.0 {
		t.Fatalf("12.0 is expected Percentage")
	}

	if rule.ChangePercentage(1.0) != 11.0 {
		t.Fatalf("11.0 is expected Percentage")
	}
}
