package trade

import "fmt"

type EnterShort struct {
	next Flow
}

func (es *EnterShort) Execute(t *Trade) {
	fmt.Println("Flow: Enter Short Position")

	if t.nxtPos == ENTER_SHORT {
		// Reset NextPosition of Short entry rules don't match
		if !t.Strategy.ShouldEnterShort(t.Series.LastIndex()) {
			t.nxtPos = ""
		}
	}

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
	fmt.Println("Flow: Exit Short Position")

	if t.nxtPos == EXIT_SHORT {
		// Validation ExitShort rules if NextPosition is ExitShort
		// If false, reset NextPosition
	}

	if es.next != nil {
		es.next.Execute(t)
	}
}

func (es *ExitShort) SetNext(next Flow) {
	es.next = next
}
