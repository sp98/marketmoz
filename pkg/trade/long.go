package trade

import "fmt"

type EnterLong struct {
	next Flow
}

func (el *EnterLong) Execute(t *Trade) {
	t.NextPosition = "EnterLong"
	fmt.Printf("%+v\n", t)

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
	t.NextPosition = "ExitLong"
	fmt.Printf("%+v\n", t)

	if el.next != nil {
		el.next.Execute(t)
	}
}

func (el *ExitLong) SetNext(next Flow) {
	el.next = next
}
