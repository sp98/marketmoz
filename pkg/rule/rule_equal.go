package rule

import "github.com/sp98/techan"

type equalRule struct {
	first  techan.Indicator
	second techan.Indicator
}

func NewEqualRule(first, second techan.Indicator) Rule {
	return equalRule{first, second}
}

func (r equalRule) IsSatisfied(index int) bool {
	return r.first.Calculate(index).EQ(r.second.Calculate(index))
}
