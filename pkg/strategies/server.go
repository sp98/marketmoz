package strategies

import (
	"fmt"

	"github.com/sp98/marketmoz/pkg/trade"
)

const (
	PVT_STRATEGY = "pvt"
)

func Start(name string) error {
	switch name {
	case PVT_STRATEGY:
		// exampleStrategy()

		exitShort := &trade.ExitShort{}
		exitLong := &trade.ExitLong{}
		exitLong.SetNext(exitShort)

		enterLong := &trade.EnterLong{}
		enterShort := &trade.EnterShort{}
		enterLong.SetNext(enterShort)
		enterShort.SetNext(exitLong)

		ohlc := trade.OHLC{}
		ohlc.SetNext(enterLong)

		t := trade.Trade{Name: "PVT"}
		ohlc.Execute(&t)

		fmt.Printf("%+v", t)

	default:
		return fmt.Errorf("invalid strategy %q", name)
	}
	return nil
}
