package trade

import "fmt"

type EnterLong struct {
	next Flow
}

func (el *EnterLong) Execute(t *Trade) {
	fmt.Println("Flow: Enter Long Position")

	if t.NextPosition == "EnterLong" {
		// update NextPosition to EnterShort if Long entry rules don't match
		if !t.Strategy.ShouldEnterLong(t.Series.LastIndex()) {
			t.NextPosition = "EnterShort"
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

	if t.NextPosition == "ExitLong" {
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
