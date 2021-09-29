package trade

import (
	"fmt"
	"os"

	"github.com/ashwanthkumar/slack-go-webhook"
	"github.com/sp98/marketmoz/pkg/common"
	"github.com/sp98/marketmoz/pkg/utils"
)

type Order struct {
	next Flow
}

func (o *Order) Execute(t *Trade) {
	fmt.Println("Flow: Create/Modify/Exit order")

	// Perform actions if NextPosition is not empty
	if t.nxtPos != "" {
		switch t.nxtPos {
		case ENTER_LONG:
		case ENTER_SHORT:
		case EXIT_LONG:
		case EXIT_SHORT:
		}

		message := notificationMessage(t)
		err := t.Notify(message)
		if err != nil {
			fmt.Printf("failed to send notification message. Error %v\n", err)
		}
	}

	if o.next != nil {
		o.next.Execute(t)
	}

	// TODO: Result trade properties like NextPosition, Stragegy, OrderParams
}

func (o *Order) SetNext(next Flow) {
	o.next = next
}

// notificationMessage sets a new notification message to be
func notificationMessage(t *Trade) slack.Payload {
	message := fmt.Sprintf(common.SLACK_NOTIFICATION_MESSAGE,
		t.nxtPos,
		t.Instrument.Name, t.Instrument.Symbol,
		t.Instrument.Exchange, t.Instrument.Segment, utils.CurrentTime())

	return slack.Payload{
		Text:    message,
		Channel: os.Getenv(common.SLACK_CHANNEL),
	}
}
