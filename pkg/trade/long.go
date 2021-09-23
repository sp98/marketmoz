package trade

import "fmt"

type EnterLong struct {
	next Flow
}

func (el *EnterLong) Execute(t *Trade) {
	fmt.Println("Flow: Enter Long Position")

	// Validation EnterLong rules if NextPosition is EnterLong
	// If false, reset NextPosition

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

	// Validation ExitLong rules if NextPosition is ExitLong
	// If false, reset NextPosition

	if el.next != nil {
		el.next.Execute(t)
	}
}

func (el *ExitLong) SetNext(next Flow) {
	el.next = next
}
