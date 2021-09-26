package data

import (
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

type Instrument struct {
	Name           string
	Symbol         string
	Token          uint32
	Exchange       string
	InstrumentType string
	Segment        string
}

type OHLC struct {
	Open   float64
	High   float64
	Low    float64
	Close  float64
	Time   int64
	Volume float64
}

func NewInstrument(name, symbol, exchange, instrumentType, segment string, token uint32) *Instrument {
	return &Instrument{
		Name:           name,
		Symbol:         symbol,
		Token:          token,
		Exchange:       exchange,
		InstrumentType: instrumentType,
		Segment:        segment,
	}
}

func (i Instrument) GetBucket(timeFrame string) string {
	return fmt.Sprintf(common.OHLC_DOWNSAMPLE_BUCKET, i.Exchange, i.Segment, timeFrame)
}

func (i Instrument) GetMeasurement() string {
	return fmt.Sprintf(common.OHLC_DOWNSAMPLE_MEASUREMENT, i.Token)
}

func (i Instrument) GetQuery(timeFrame string) (string, error) {
	queryBytes, err := assets.ReadFileAndReplace(
		ohlcQueryAsset,
		[]string{
			"${INPUT_BUCKET}", i.GetBucket(timeFrame),
			"${INPUT_MEASUREMENT}", i.GetMeasurement(),
			"${INPUT_EVERY}", timeFrame,
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

func (i Instrument) GetOHLC(db *influx.DB) (*[]OHLC, error) {
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
	return ohlc, nil
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
