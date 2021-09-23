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

		order := &trade.Order{}

		exitShort := &trade.ExitShort{}
		exitShort.SetNext(order)
		exitLong := &trade.ExitLong{}
		exitLong.SetNext(exitShort)

		enterLong := &trade.EnterLong{}
		enterShort := &trade.EnterShort{}
		enterLong.SetNext(enterShort)
		enterShort.SetNext(exitLong)

		position := &trade.Position{}
		position.SetNext(enterLong)

		rules := trade.Rules{}
		rules.SetNext(position)

		t := &trade.Trade{Name: "PVT"}
		rules.Execute(t)

		fmt.Printf("%+v \n", t)

	default:
		return fmt.Errorf("invalid strategy %q", name)
	}
	return nil
}
