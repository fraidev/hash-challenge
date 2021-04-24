package services

import (
	"time"

	models "github.com/fraifelipe/hash-challenge/backend/discounts/models"
	pb "github.com/fraifelipe/hash-challenge/backend/discounts/protos/discount"
)

func CalculateDiscount(product *models.Product, user *models.User) pb.ProductDiscount {
	var percentageDiscount float32 = 0.0

	if IsUserBirthDate(user) {
		percentageDiscount += 5.0
	}

	if IsBlackFriday() {
		percentageDiscount += 10.0
	}

	if percentageDiscount > 10.0 {
		percentageDiscount = 10.0
	}

	//calculte discount value
	valueInCents := int32(float32(product.PriceInCents/100) * percentageDiscount)

	return pb.ProductDiscount{
		Percentage:   percentageDiscount,
		ValueInCents: valueInCents,
		ProductId:    product.ID.String(),
	}
}

func IsThisDayOfYear(month int, day int) bool {
	currentDate := time.Now()
	return month == int(currentDate.Month()) && day == currentDate.Day()
}

func IsUserBirthDate(user *models.User) bool {
	return IsThisDayOfYear(int(user.DateOfBirth.Month()), user.DateOfBirth.Day())
}

func IsBlackFriday() bool {
	return IsThisDayOfYear(11, 25)
}
