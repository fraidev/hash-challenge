package services

type BlackFridayRule struct {
}

func (_ BlackFridayRule) Validate(ruleEngine *RuleEngine) bool {
	return IsToday(11, 25)
}

func (_ BlackFridayRule) ChangePercentage(percentageDiscount float32) float32 {
	percentageDiscount += 10.0
	return percentageDiscount
}
