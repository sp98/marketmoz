package trade

import (
	"fmt"
)

type OHLC struct {
	next Flow
}

func (o *OHLC) Execute(t *Trade) {
	t.NextPosition = "GETOHLC"
	fmt.Printf("%+v\n", t)
	if o.next != nil {
		o.next.Execute(t)
	}
}

func (o *OHLC) SetNext(next Flow) {
	o.next = next
}
