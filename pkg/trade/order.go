package trade

import (
	"fmt"
	"os"

	"github.com/ashwanthkumar/slack-go-webhook"
	"github.com/sp98/marketmoz/pkg/common"
	"github.com/sp98/marketmoz/pkg/utils"
	kiteconnect "github.com/zerodha/gokiteconnect/v4"
	"go.uber.org/zap"
)

type Order struct {
	next Flow
}

func (o *Order) Execute(t *Trade) {
	Logger.Info("Flow: Create/Modify/Exit order")

	// Reset trade properties like NextPosition and OrderParams at the end
	defer func() {
		t.ResetNextPosition()
		t.ResetOrderParams()
	}()

	var res kiteconnect.OrderResponse
	var err error
	var status string
	// Perform actions if NextPosition is not empty
	if t.nxtPos != "" {
		switch t.nxtPos {
		case ENTER_LONG, ENTER_SHORT:
			res, err = t.KClient.PlaceOrder("regular", *t.OrderParams)
		case EXIT_LONG, EXIT_SHORT:
			//t.KClient.ExitOrder("regular", "", "")
		}

		if err != nil {
			Logger.Error("failed to execute order", zap.Any("order", t.nxtPos), zap.Error(err))
			status = "FAILURE"
		} else {
			Logger.Info("Successfully executed order", zap.Any("order", t.nxtPos), zap.Any("response", res))
			status = "SUCCESS"
		}

		message := notificationMessage(t, status)
		err := t.Notify(message)
		if err != nil {
			Logger.Error("failed to send notification message", zap.Errors("error", err))
			return
		}
	}

	if o.next != nil {
		o.next.Execute(t)
	}
}

func (o *Order) SetNext(next Flow) {
	o.next = next
}

// notificationMessage sets a new notification message to be
func notificationMessage(t *Trade, status string) slack.Payload {
	message := fmt.Sprintf(common.SLACK_NOTIFICATION_MESSAGE,
		t.nxtPos,
		status,
		t.Instrument.Name, t.Instrument.Symbol,
		t.Instrument.Exchange, t.Instrument.Segment, utils.CurrentTime())

	return slack.Payload{
		Text:    message,
		Channel: os.Getenv(common.SLACK_CHANNEL),
	}
}
