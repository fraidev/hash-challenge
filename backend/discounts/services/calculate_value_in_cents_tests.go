package services

import (
	"testing"
)

func TestCalculateValueInCents(t *testing.T) {
	if CalculateValueInCents(50000, 10) != 5000 {
		t.Fatalf("5000 is expected Percentage")
	}
	if CalculateValueInCents(25000, 10) != 2500 {
		t.Fatalf("2500 is expected Percentage")
	}

	if CalculateValueInCents(10000, 5) != 500 {
		t.Fatalf("500 is expected Percentage")
	}

	if CalculateValueInCents(1000, 1) != 10 {
		t.Fatalf("10 is expected Percentage")
	}
}
