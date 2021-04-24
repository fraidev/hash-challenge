package services

import (
	"testing"
	"time"

	models "github.com/fraifelipe/hash-challenge/backend/discounts/models"
	uuid "github.com/satori/go.uuid"
)

func TestCalculateDiscountBirthDate(t *testing.T) {
	product := models.Product{
		ID:           uuid.NewV1(),
		PriceInCents: 10000,
	}
	user := models.User{
		DateOfBirth: time.Now(),
	}

	result := CalculateDiscount(&product, &user)

	if result.Percentage != 5 || result.ValueInCents != 500 || result.ProductId != product.ID.String() && !IsToday(11, 25) {
		t.Fatalf("Result not expected: %+v", &result)
	}

}

func TestCalculateDiscount(t *testing.T) {
	product := models.Product{
		ID:           uuid.NewV1(),
		PriceInCents: 10000,
	}
	user := models.User{
		DateOfBirth: time.Now().AddDate(0, 0, -1),
	}

	result := CalculateDiscount(&product, &user)

	if result.Percentage != 0 || result.ValueInCents != 0 || result.ProductId != product.ID.String() && !IsToday(11, 25) {
		t.Fatalf("Result not expected: %+v", &result)
	}
}
