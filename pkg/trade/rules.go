package trade

import (
	"github.com/sp98/marketmoz/pkg/strategy"
)

type Rules struct {
	next Flow
}

func (r *Rules) Execute(t *Trade) error {
	Logger.Info("Flow: Set Rules")
	t.SetStrategy(strategy.PVTStrategyRules(t.Series))

	if r.next != nil {
		return r.next.Execute(t)
	}

	return nil
}

func (r *Rules) SetNext(next Flow) {
	r.next = next
}
