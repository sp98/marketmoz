package trade

import (
	"fmt"

	"github.com/sp98/marketmoz/pkg/common"
	"github.com/sp98/marketmoz/pkg/strategy"
	kiteconnect "github.com/zerodha/gokiteconnect/v4"
	"go.uber.org/zap"
)

type EnterLong struct {
	next Flow
}

func (el *EnterLong) Execute(t *Trade) {
	Logger.Info("Flow: Enter Long Position")

	if t.nxtPos == ENTER_LONG {
		if t.Strategy.ShouldEnterLong(t.Series.LastIndex()) {
			// Find R2R ratio
			triggerPrice := strategy.GetPVTStrategyLongSL(t.Series)
			query, err := t.Instrument.GetQuery("5m", common.LASTPRICE_QUERY_ASSET)
			if err != nil {
				Logger.Error("failed to get query for last price", zap.Error(err))
				return
			}

			// How to ensure that last Price data is from the last time frame
			lastPrice, err := t.Instrument.GetLastPrice(t.DB, query)
			if err != nil {
				Logger.Error("failed to to get last price", zap.Error(err))
			}

			if triggerPrice > lastPrice {
				Logger.Warn("trigger price can't be higher than last price for next trade", zap.Any("nextTrade", t.nxtPos))
				return
			}
			risk := lastPrice - triggerPrice
			reward := lastPrice + (2 * risk) // 1:2
			Logger.Info("Risk:Reward ", zap.Float64("risk", risk), zap.Float64("reward", reward))

			// TODO: Ignore if risk is too high?

			// Create order for long position
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
				TriggerPrice:      triggerPrice, // Lower of Previous candle and current Candle Low
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
