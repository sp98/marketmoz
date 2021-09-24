package trade

import (
	"fmt"

	"github.com/sp98/techan"
)

type Series struct {
	next Flow
}

func (s *Series) Execute(t *Trade) {
	fmt.Println("Flow: Get Time Series data")
	// Get OHLC data and build the strategy rules
	series := techan.TimeSeries{}
	t.Series = &series

	if s.next != nil {
		s.next.Execute(t)
	}
}

func (s *Series) SetNext(next Flow) {
	s.next = next
}
