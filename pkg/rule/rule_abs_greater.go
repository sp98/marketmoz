package rule

import "github.com/sp98/techan"

type absGreaterRule struct {
	first  techan.Indicator
	second techan.Indicator
}

func NewAbsGreaterRule(first, second techan.Indicator) Rule {
	return absGreaterRule{first, second}
}

func (r absGreaterRule) IsSatisfied(index int) bool {
	return r.first.Calculate(index).Abs().GTE(r.second.Calculate(index))
}
