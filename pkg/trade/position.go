package trade

import "fmt"

type Position struct {
	next Flow
}

func (p *Position) Execute(t *Trade) error {
	fmt.Println("Flow: Get Position")

	orders, err := t.KClient.GetOrders()
	if err != nil {
		return fmt.Errorf("failed to get existing orders. Error %v", err)
	}

	// Set next position to ENTER_LONG
	t.nxtPos = ENTER_LONG
	for _, order := range orders {
		if order.InstrumentToken == t.Instrument.Token {
			if order.Status == "OPEN" {
				if order.TransactionType == "BUY" {
					t.nxtPos = EXIT_LONG
				} else if order.TransactionType == "SELL" {
					t.nxtPos = EXIT_SHORT
				}
			}
		}
	}

	if p.next != nil {
		return p.next.Execute(t)
	}

	return nil
}

func (p *Position) SetNext(next Flow) {
	p.next = next
}
