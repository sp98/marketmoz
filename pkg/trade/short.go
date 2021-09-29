package trade

import (
	"fmt"

	"github.com/sp98/marketmoz/pkg/strategy"
	kiteconnect "github.com/zerodha/gokiteconnect/v4"
)

type EnterShort struct {
	next Flow
}

func (es *EnterShort) Execute(t *Trade) {
	fmt.Println("Flow: Enter Short Position")

	if t.nxtPos == ENTER_SHORT {
		if t.Strategy.ShouldEnterShort(t.Series.LastIndex()) {
			// Create order for long position
			triggerPrice := strategy.GetPVTStrategyShortSL(t.Series)
			orderParams := &kiteconnect.OrderParams{
				Exchange:          t.Instrument.Exchange,
				Tradingsymbol:     t.Instrument.Symbol,
				Validity:          "DAY",
				Product:           "MIS",
				OrderType:         "SL-M",
				TransactionType:   "SELL",
				Quantity:          0,
				DisclosedQuantity: 0,
				Price:             0,
				TriggerPrice:      triggerPrice.Float(), // Higher of Previous candle and current Candle High
				Squareoff:         0,
				Stoploss:          0,
				TrailingStoploss:  0,
				Tag:               t.Name,
			}
			t.SetOrderParams(orderParams)
		} else {
			// Reset NextPosition of Short entry rules don't match
			t.nxtPos = ""
		}
	}

	if es.next != nil {
		es.next.Execute(t)
	}
}

func (es *EnterShort) SetNext(next Flow) {
	es.next = next
}

type ExitShort struct {
	next Flow
}

func (es *ExitShort) Execute(t *Trade) {
	fmt.Println("Flow: Exit Short Position")

	if t.nxtPos == EXIT_SHORT {
		// Validation ExitShort rules if NextPosition is ExitShort
		// If false, reset NextPosition
	}

	if es.next != nil {
		es.next.Execute(t)
	}
}

func (es *ExitShort) SetNext(next Flow) {
	es.next = next
}
