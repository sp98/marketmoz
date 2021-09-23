package trade

import (
	"fmt"

	"github.com/sp98/marketmoz/pkg/strategy"
	"github.com/sp98/techan"
)

type Rules struct {
	next Flow
}

func (r *Rules) Execute(t *Trade) {
	fmt.Println("Flow: Set Rules")
	// Get OHLC data and build the strategy rules
	series := techan.TimeSeries{}
	t.Strategy = strategy.PVTStrategyRules(&series)

	if r.next != nil {
		r.next.Execute(t)
	}
}

func (r *Rules) SetNext(next Flow) {
	r.next = next
}
