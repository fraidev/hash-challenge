package services

type MaxPercentageRule struct {
}

func (_ MaxPercentageRule) Validate(ruleEngine *RuleEngine) bool {
	return ruleEngine.percentageDiscount > 10.0
}

func (_ MaxPercentageRule) ChangePercentage(percentageDiscount float32) float32 {
	percentageDiscount = 10.0
	return percentageDiscount
}
