package services

type Rule interface {
	Validade(ruleEngine *RuleEngine) bool
	ChangePercentage(percentageDiscount float32) float32
}

type RuleEngine struct {
	rules              []Rule
	percentageDiscount float32
}

func NewRuleEngine() *RuleEngine {
	return &RuleEngine{}
}

func (ruleEngine *RuleEngine) ApppendRule(rule Rule) {
	ruleEngine.rules = append(ruleEngine.rules, rule)
}

func (ruleEngine *RuleEngine) CalculateDiscount() {
	for _, rule := range ruleEngine.rules {
		if rule.Validade(ruleEngine) {
			ruleEngine.percentageDiscount = rule.ChangePercentage(ruleEngine.percentageDiscount)
		}
	}
}
