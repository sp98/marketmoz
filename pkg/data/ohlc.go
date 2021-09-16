package data

import (
	"context"
	"fmt"

	"github.com/influxdata/influxdb-client-go/v2/api"
	ms "github.com/mitchellh/mapstructure"
	"github.com/sp98/marketmoz/assets"
	"github.com/sp98/marketmoz/pkg/common"
	"github.com/sp98/marketmoz/pkg/db/influx"
	"go.uber.org/zap"
)

var (
	Logger              *zap.Logger
	ohlcQueryAsset      = "queries/ohlc.flux"
	ohlcQueryTestAssert = "queries/test.flux"
)

type OHLCData struct {
	org      string
	token    string
	cadence  string
	exchange string
	segment  string
	data     *[]OHLC
}

type OHLC struct {
	Open   float64
	High   float64
	Low    float64
	Close  float64
	Time   int64
	Volume float64
}

func NewOHLCData(org, token, cadence, exchange, segment string) *OHLCData {
	return &OHLCData{
		org:      org,
		token:    token,
		cadence:  cadence,
		segment:  segment,
		exchange: exchange,
	}
}

func (o OHLCData) SetOrg(org string) OHLCData {
	o.org = org
	return o
}

func (o OHLCData) SetToken(token string) OHLCData {
	o.token = token
	return o
}

func (o OHLCData) SetCadence(cadence string) OHLCData {
	o.cadence = cadence
	return o
}

func (o OHLCData) SetExchange(exchange string) OHLCData {
	o.exchange = exchange
	return o
}

func (o OHLCData) SetSegment(segment string) OHLCData {
	o.segment = segment
	return o
}

func (o OHLCData) GetData() *[]OHLC {
	return o.data
}

func (o OHLCData) GetBucket() string {
	return fmt.Sprintf(common.OHLC_DOWNSAMPLE_BUCKET, o.exchange, o.segment, o.cadence)
}

func (o OHLCData) GetMeasurement() string {
	return fmt.Sprintf(common.OHLC_DOWNSAMPLE_MEASUREMENT, o.token)
}

func (o OHLCData) GetQuery() (string, error) {
	queryBytes, err := assets.ReadFileAndReplace(
		ohlcQueryAsset,
		[]string{
			"${INPUT_BUCKET}", o.GetBucket(),
			"${INPUT_MEASUREMENT}", o.GetMeasurement(),
			"${INPUT_EVERY}", o.cadence,
		},
	)

	if err != nil {
		return "", err
	}

	return string(queryBytes), nil
}

func getTestQuery() (string, error) {
	queryBytes, err := assets.ReadFileAndReplace(
		ohlcQueryTestAssert,
		[]string{},
	)
	return string(queryBytes), err
}

func (o OHLCData) GetOHLC() (*OHLCData, error) {
	ctx := context.Background()
	db := influx.NewDB(ctx, o.org, common.INFLUXDB_URL, common.INFLUXDB_TOKEN)
	defer db.Client.Close()

	// TODO: remove getTestQuery method when calling the kite API
	query, err := getTestQuery()
	if err != nil {
		return nil, err
	}

	result, err := db.GetData(query)
	if err != nil {
		Logger.Error("failed to get ohlc data from the influx db.", zap.Error(err))
		return nil, fmt.Errorf("failed to get ohlc data from the influx db. Error %v", err)
	}

	ohlc, err := parseOHLC(result)
	if err != nil {
		return nil, fmt.Errorf("failed to parse the results. Error %v", err)

	}

	Logger.Info("OHCL result ", zap.Any("ohlc", ohlc))
	o.data = ohlc
	return &o, nil
}

func parseOHLC(in *api.QueryTableResult) (*[]OHLC, error) {
	out, err := parse(in)
	if err != nil {
		return nil, fmt.Errorf("failed to parse the results. Error %v", err)

	}
	ohlc := &[]OHLC{}
	err = ms.Decode(out, ohlc)
	if err != nil {
		Logger.Error("failed to decode parsed db data into ohlc struct", zap.Error(err))
		return nil, err
	}

	return ohlc, nil
}
