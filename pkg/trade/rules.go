package trade

import (
	"fmt"

	"github.com/sp98/marketmoz/pkg/strategy"
)

type Rules struct {
	next Flow
}

func (r *Rules) Execute(t *Trade) {
	fmt.Println("Flow: Set Rules")
	t.SetStrategy(strategy.PVTStrategyRules(t.Series))

	if r.next != nil {
		r.next.Execute(t)
	}
}

func (r *Rules) SetNext(next Flow) {
	r.next = next
}
