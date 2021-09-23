package trade

import "fmt"

type EnterShort struct {
	next Flow
}

func (es *EnterShort) Execute(t *Trade) {
	t.NextPosition = "EnterShort"
	fmt.Printf("%+v\n", t)

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
	t.NextPosition = "ExitShort"
	fmt.Printf("%+v\n", t)

	if es.next != nil {
		es.next.Execute(t)
	}
}

func (es *ExitShort) SetNext(next Flow) {
	es.next = next
}
