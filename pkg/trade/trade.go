package trade

import (
	"github.com/sp98/marketmoz/pkg/db/influx"
	"github.com/sp98/marketmoz/pkg/strategy"
	kiteconnect "github.com/zerodha/gokiteconnect/v4"
)

type Trade struct {
	// Name of the strategy to be used
	Name string

	// NextPosition defines the next set of actions to be taken
	NextPosition string

	// Strategy defines the set of rules to be used in this strategy
	Strategy strategy.Strategy

	// KClient is the client to call the Zerodha API
	KClient *kiteconnect.Client

	// DB is the influx DB to get OHLC data
	DB *influx.DB
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

func (t *Trade) SetNextPosition(position string) {
	t.NextPosition = position
}

func (t *Trade) SetStrategy(strategy strategy.Strategy) {
	t.Strategy = strategy
}
