package services

import models "github.com/fraifelipe/hash-challenge/backend/discounts/models"

type BirthDateRule struct {
	user *models.User
}

func (birthDateRule BirthDateRule) Validate(ruleEngine *RuleEngine) bool {
	return IsToday(int(birthDateRule.user.DateOfBirth.Month()), birthDateRule.user.DateOfBirth.Day())
}

func (birthDateRule BirthDateRule) ChangePercentage(percentageDiscount float32) float32 {
	percentageDiscount += 5.0
	return percentageDiscount
}
