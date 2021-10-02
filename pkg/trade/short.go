package trade

import (
	"fmt"

	"github.com/sp98/marketmoz/pkg/common"
	"github.com/sp98/marketmoz/pkg/strategy"
	kiteconnect "github.com/zerodha/gokiteconnect/v4"
	"go.uber.org/zap"
)

type EnterShort struct {
	next Flow
}

func (es *EnterShort) Execute(t *Trade) {
	fmt.Println("Flow: Enter Short Position")

	if t.nxtPos == ENTER_SHORT {
		if t.Strategy.ShouldEnterShort(t.Series.LastIndex()) {
			// Find R2R ratio
			triggerPrice := strategy.GetPVTStrategyShortSL(t.Series)
			query, err := t.Instrument.GetQuery(t.Interval, common.LASTPRICE_QUERY_ASSET)
			if err != nil {
				Logger.Error("failed to get query for last price", zap.Error(err))
				return
			}

			// How to ensure that last Price data is from the last time frame
			lastPrice, err := t.Instrument.GetLastPrice(t.DB, query)
			if err != nil {
				Logger.Error("failed to to get last price", zap.Error(err))
			}

			if triggerPrice < lastPrice {
				Logger.Warn("trigger price can't be lower than last price for trade", zap.Any("nextPosition", t.nxtPos))
				return
			}
			risk := triggerPrice - lastPrice
			reward := lastPrice - (2 * risk) // 1:2
			Logger.Info("Risk:Reward ", zap.Float64("risk", risk), zap.Float64("reward", reward))

			// Create order for long position
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
				TriggerPrice:      triggerPrice, // Higher of Previous candle and current Candle High
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
