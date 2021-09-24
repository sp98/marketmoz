package kite

import (
	"fmt"
	"time"

	"github.com/sp98/marketmoz/pkg/common"
	"github.com/sp98/marketmoz/pkg/db/influx"
	kiteconnect "github.com/zerodha/gokiteconnect/v4"
	kitemodels "github.com/zerodha/gokiteconnect/v4/models"
	kiteticker "github.com/zerodha/gokiteconnect/v4/ticker"
	"go.uber.org/zap"
)

var Logger *zap.Logger

type Kite struct {
	// Client is the KiteConnect client
	Client *kiteconnect.Client

	// TClient is the client for streaming ticks.
	TClient *kiteticker.Ticker

	// User session
	User *kiteconnect.UserSession

	// Subscriptions
	Subscriptions []uint32

	// Store represents the underlying storage for tick data
	Store *influx.DB
}

func New(apiKey, requestToken string, subs []uint32) (*Kite, error) {
	c, user, err := NewKiteConnectClient(apiKey, requestToken)
	if err != nil {
		return nil, err
	}

	tc := NewTickerClient(apiKey, user.AccessToken)

	return &Kite{
		Client:        c,
		User:          user,
		TClient:       tc,
		Subscriptions: subs,
	}, nil

}

func NewKiteConnectClient(apiSecret, requestToken string) (*kiteconnect.Client, *kiteconnect.UserSession, error) {
	kc := kiteconnect.New(apiSecret)
	// Get user details and access token
	data, err := kc.GenerateSession(requestToken, apiSecret)
	if err != nil {
		return nil, nil, err
	}

	kc.SetAccessToken(data.AccessToken)

	return kc, &data, nil
}

func NewTickerClient(apiKey, accessToken string) *kiteticker.Ticker {
	return kiteticker.New(apiKey, accessToken)
}

func (k *Kite) StartKiteFetcher() {
	// Assign callbacks
	k.onConnect()
	k.onReconnect()
	k.onTick()
	k.onError()
	k.onClose()
	k.onOrderUpdate()

	// Start the connection
	k.TClient.Serve()
}

func (k *Kite) onTick() {
	onTick := func(tick kitemodels.Tick) {
		Logger.Info("tick received", zap.Any("tick", tick))
		k.storeTick(tick)
	}

	k.TClient.OnTick(onTick)
}

func (k *Kite) onConnect() {
	onConnect := func() {
		Logger.Info("connected to kite successfully")
		err := k.TClient.Subscribe(k.Subscriptions)
		if err != nil {
			Logger.Error("failed to add subscriptions", zap.Error(err))
		}
	}
	k.TClient.OnConnect(onConnect)
}

func (k *Kite) onReconnect() {
	onReconnect := func(attempt int, delay time.Duration) {
		Logger.Info("attempting to reconnect", zap.Int("attempt", attempt), zap.Duration("duration", delay))
	}
	k.TClient.OnReconnect(onReconnect)
}

func (k *Kite) onError() {
	onError := func(err error) {
		Logger.Error("failed to fetch tick data", zap.Error(err))
	}
	k.TClient.OnError(onError)
}

func (k *Kite) onClose() {
	onClose := func(code int, reason string) {
		Logger.Info("kite websocket connection closed", zap.Int("code", code), zap.String("reason", reason))
	}
	k.TClient.OnClose(onClose)
}

func (k *Kite) onOrderUpdate() {
	onOrderUpdate := func(order kiteconnect.Order) {
		Logger.Info("order updated", zap.String("orderID", order.OrderID))
	}
	k.TClient.OnOrderUpdate(onOrderUpdate)
}

func (k *Kite) storeTick(tick kitemodels.Tick) {
	tags := map[string]string{}
	bucket, err := getRTDBucket(fmt.Sprintf("%d", tick.InstrumentToken))
	if err != nil {
		Logger.Error("failed to get real time data bucket name", zap.Uint32("token", tick.InstrumentToken))
		return
	}
	measurement := fmt.Sprintf(common.REAL_TIME_DATA_MEASUREMENT, tick.InstrumentToken)

	fields := map[string]interface{}{
		"LastPrice": tick.LastPrice,
		"Volume":    tick.VolumeTraded,
	}

	err = k.Store.WriteData(bucket, measurement, tags, fields, tick.Timestamp.Time)
	if err != nil {
		Logger.Error("failed to write real time data", zap.Uint32("token", tick.InstrumentToken), zap.Error(err))
	}
}

func (k *Kite) CreateDownsampleTasks() error {
	tasks, err := GetOHLCDownSamplingTasks()
	if err != nil {
		return err
	}
	for _, task := range *tasks {
		_, err := k.Store.WriteTask(&task)
		if err != nil {
			Logger.Error("failed to create task", zap.String("taskname", task.Name), zap.Error(err))
			continue
		}
	}

	return nil
}

// getRTDBucket returns bucket name to store real time data.
func getRTDBucket(token string) (string, error) {
	td := common.GetInstrumentDetails(token)
	if td == nil {
		return "", fmt.Errorf("failed to get token details for token %q", token)
	}

	b := fmt.Sprintf(common.REAL_TIME_DATA_BUCKET, td.Exchange, td.Segment)
	return b, nil

}
