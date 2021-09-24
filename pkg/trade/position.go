package trade

import "fmt"

type Position struct {
	next Flow
}

func (p *Position) Execute(t *Trade) {
	fmt.Println("Flow: Get Position")

	/*
		TODO: get current position of the insturment using the broker client.
		Set NextPosition to:
			EnterLong:  If no position found.
			ExitLong: If Buy position exists
			ExitShort: If Sell Position exits.
	*/

	if p.next != nil {
		p.next.Execute(t)
	}
}

func (p *Position) SetNext(next Flow) {
	p.next = next
}
