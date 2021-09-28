package trade

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/ashwanthkumar/slack-go-webhook"
	"github.com/sp98/marketmoz/pkg/common"
	"github.com/sp98/marketmoz/pkg/data"
	"github.com/sp98/marketmoz/pkg/db/influx"
	"github.com/sp98/marketmoz/pkg/fetcher/kite"
	"github.com/sp98/marketmoz/pkg/strategy"
	"github.com/sp98/marketmoz/pkg/utils"
	"github.com/sp98/techan"
	kiteconnect "github.com/zerodha/gokiteconnect/v4"
)

type NextPosition string

const (
	ENTER_LONG  NextPosition = "ENTER_LONG"
	EXIT_LONG   NextPosition = "EXIT_LONG"
	ENTER_SHORT NextPosition = "ENTER_SHORT"
	EXIT_SHORT  NextPosition = "EXIT_SHORT"
)

const (
	PVT_STRATEGY = "pvt"
)

type Trade struct {
	// Name of the strategy to be used
	Name string

	Series *techan.TimeSeries

	// NextPosition defines the next set of actions to be taken
	nxtPos NextPosition

	// Strategy defines the set of rules to be used in this strategy
	Strategy strategy.Strategy

	// KClient is the client to call the Zerodha API
	KClient *kiteconnect.Client

	// DB is the influx DB to get OHLC data
	DB *influx.DB

	// Instrument to be traded
	Instrument data.Instrument

	// OrderParams represents the parameters for the new long or short order
	OrderParams kiteconnect.OrderParams
}

func NewTrade(name string) *Trade {
	return &Trade{Name: name}
}

func (t *Trade) SetBrokerClient(client *kiteconnect.Client) {
	t.KClient = client
}

func (t *Trade) SetDB(db *influx.DB) {
	t.DB = db
}

func (t *Trade) SetNextPosition(position NextPosition) {
	t.nxtPos = position
}

func (t *Trade) ResetNextPosition() {
	t.nxtPos = ""
}

func (t *Trade) SetStrategy(strategy strategy.Strategy) {
	t.Strategy = strategy
}

func (t *Trade) SetInstrument(instrument data.Instrument) {
	t.Instrument = instrument
}

func (t *Trade) SetOrderParams(orderParams kiteconnect.OrderParams) {
	t.OrderParams = orderParams
}

func (t *Trade) Notify(payload slack.Payload) []error {
	url := os.Getenv(common.SLACK_WEBHOOK_URL)
	err := slack.Send(url, "", payload)
	if err != nil {
		return err
	}
	return nil
}

func (t *Trade) start(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	start, _ := utils.StartTimeAndLoc()
	flow := NewTradeFlow()
	for range utils.Every(ctx, start, 1*time.Minute) {
		flow.Execute(t)
		fmt.Printf("Trade - %+v\n", t)
	}
}

func Start(name string) error {
	// Get kite connect client
	// strategy.ExampleStrategy()
	apiKey := os.Getenv(common.KITE_API_KEY)
	apiSecret := os.Getenv(common.KITE_API_SECRET)
	requestToken := os.Getenv(common.KITE_REQUEST_TOKEN)

	client, _, err := kite.NewKiteConnectClient(apiKey, apiSecret, requestToken)
	if err != nil {
		//return fmt.Errorf("failed to get kite client. Error %v", err)
	}

	// Get influx DB Client
	db := influx.NewDB(context.TODO(), common.INFLUXDB_ORGANIZATION,
		common.INFLUXDB_URL, common.INFLUXDB_TOKEN)

	// Setup trade flow

	ctx, _ := context.WithDeadline(context.Background(), utils.EndTime())
	switch name {
	case PVT_STRATEGY:
		var wg sync.WaitGroup
		pvtInstruments := strategy.GetPVTInstruments()
		for _, instrument := range *pvtInstruments {
			wg.Add(1)
			t := NewTrade("PVT")
			t.SetDB(db)
			t.SetBrokerClient(client)
			t.SetInstrument(instrument)
			go t.start(ctx, &wg)
		}
		wg.Wait()

	default:
		return fmt.Errorf("invalid strategy %q", name)
	}

	fmt.Println("Trading with PVT strategy ended at - ", utils.CurrentTime())

	// TODO: Clean up after trade ended.
	// 1. Exit any pending trade.
	// 2. Send notification to end any pending trade
	return nil
}

func NewTradeFlow() Flow {
	order := &Order{}

	exitShort := &ExitShort{}
	exitShort.SetNext(order)
	exitLong := &ExitLong{}
	exitLong.SetNext(exitShort)

	enterLong := &EnterLong{}
	enterShort := &EnterShort{}
	enterLong.SetNext(enterShort)
	enterShort.SetNext(exitLong)

	position := &Position{}
	position.SetNext(enterLong)

	rules := &Rules{}
	rules.SetNext(position)

	series := &Series{}
	series.SetNext(rules)

	return series
}
