package services

import (
	models "github.com/fraifelipe/hash-challenge/backend/discounts/models"
	pb "github.com/fraifelipe/hash-challenge/backend/discounts/protos/discount"
)

func CalculateDiscount(product *models.Product, user *models.User) pb.ProductDiscount {

	ruleEngine := NewRuleEngine()

	ruleEngine.ApppendRule(BirthDateRule{user: user})
	ruleEngine.ApppendRule(BlackFridayRule{})
	ruleEngine.ApppendRule(MaxPercentageRule{})

	ruleEngine.CalculateDiscount()

	//calculte discount value
	valueInCents := CalculateValueInCents(product.PriceInCents, ruleEngine.percentageDiscount)

	return pb.ProductDiscount{
		Percentage:   ruleEngine.percentageDiscount,
		ValueInCents: valueInCents,
		ProductId:    product.ID.String(),
	}
}
