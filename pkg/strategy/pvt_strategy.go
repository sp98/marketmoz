package strategy

import (
	"github.com/sp98/marketmoz/pkg/rule"
	"github.com/sp98/techan"
)

func PVTStrategyRules(series *techan.TimeSeries) Strategy {
	// Create RuleStrategy using all entry and exit rules for PVT strategy
	entryConstant := techan.NewConstantIndicator(30)
	exitConstant := techan.NewConstantIndicator(10)

	longEntryRule := rule.NewAndRule(
		rule.NewCrossDownIndicatorRule(entryConstant, exitConstant),
	)

	return &RuleStrategy{
		LongEntryRule: longEntryRule,
	}
}
