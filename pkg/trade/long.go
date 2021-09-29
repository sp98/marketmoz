package trade

import (
	"fmt"

	"github.com/sp98/marketmoz/pkg/strategy"
	kiteconnect "github.com/zerodha/gokiteconnect/v4"
)

type EnterLong struct {
	next Flow
}

func (el *EnterLong) Execute(t *Trade) {
	fmt.Println("Flow: Enter Long Position")

	if t.nxtPos == ENTER_LONG {
		if t.Strategy.ShouldEnterLong(t.Series.LastIndex()) {
			// Create order for long position
			triggerPrice := strategy.GetPVTStrategyLongSL(t.Series)
			orderParams := &kiteconnect.OrderParams{
				Exchange:          t.Instrument.Exchange,
				Tradingsymbol:     t.Instrument.Symbol,
				Validity:          "DAY",
				Product:           "MIS",
				OrderType:         "SL-M",
				TransactionType:   "BUY",
				Quantity:          0,
				DisclosedQuantity: 0,
				Price:             0,
				TriggerPrice:      triggerPrice.Float(), // Lower of Previous candle and current Candle Low
				Squareoff:         0,
				Stoploss:          0,
				TrailingStoploss:  0,
				Tag:               t.Name,
			}

			t.SetOrderParams(orderParams)
		} else {
			// update NextPosition to EnterShort if Long entry rules don't match
			t.nxtPos = ENTER_SHORT
		}
	}
	if el.next != nil {
		el.next.Execute(t)
	}
}

func (el *EnterLong) SetNext(next Flow) {
	el.next = next
}

type ExitLong struct {
	next Flow
}

func (el *ExitLong) Execute(t *Trade) {
	fmt.Println("Flow: Exit Long Position")

	if t.nxtPos == EXIT_LONG {
		// Validation ExitLong rules if NextPosition is ExitLong
		// If false, reset NextPosition
	}

	if el.next != nil {
		el.next.Execute(t)
	}
}

func (el *ExitLong) SetNext(next Flow) {
	el.next = next
}
