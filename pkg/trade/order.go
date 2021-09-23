package trade

import "fmt"

type Order struct {
	next Flow
}

func (o *Order) Execute(t *Trade) {
	fmt.Println("Flow: Create/Modify/Exit order")

	if o.next != nil {
		o.next.Execute(t)
	}
}

func (o *Order) SetNext(next Flow) {
	o.next = next
}
