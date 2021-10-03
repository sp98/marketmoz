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
	Logger *zap.Logger
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

type LastPrice struct {
	Price float64
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

// GetDSBucket returns bucket for downsampled data
func (i Instrument) GetDSBucket(timeFrame string) string {
	return fmt.Sprintf(common.OHLC_DOWNSAMPLE_BUCKET, i.InstrumentType, i.Segment, i.Exchange, timeFrame)
}

// GetDSMeasurement returns measurement for downsampled data
func (i Instrument) GetDSMeasurement() string {
	return fmt.Sprintf(common.OHLC_DOWNSAMPLE_MEASUREMENT, i.Token)
}

// GetDSQuery returns flux query for downsampled buckets
func (i Instrument) GetDSQuery(interval, queryFile string) (string, error) {
	queryBytes, err := assets.ReadFileAndReplace(
		queryFile,
		[]string{
			"${INPUT_BUCKET}", i.GetDSBucket(interval),
			"${INPUT_MEASUREMENT}", i.GetDSMeasurement(),
			"${INPUT_EVERY}", interval,
		},
	)

	if err != nil {
		return "", err
	}

	return string(queryBytes), nil
}

// GetRTDBucket returns bucket for real time data
func (i Instrument) GetRTDBucket() string {
	return fmt.Sprintf(common.REAL_TIME_DATA_BUCKET, i.InstrumentType, i.Segment, i.Exchange)
}

// GetRTDMeasurement returns measurement for real time data
func (i Instrument) GetRTDMeasurement() string {
	return fmt.Sprintf(common.REAL_TIME_DATA_MEASUREMENT, i.Token)
}

// GetRTDQuery returns flux query to get real time data
func (i Instrument) GetRTDQuery(interval, queryFile string) (string, error) {
	queryBytes, err := assets.ReadFileAndReplace(
		queryFile,
		[]string{
			"${INPUT_BUCKET}", i.GetRTDBucket(),
			"${INPUT_MEASUREMENT}", i.GetRTDMeasurement(),
			"${INPUT_EVERY}", interval,
		},
	)

	if err != nil {
		return "", err
	}

	return string(queryBytes), nil
}

func (i Instrument) GetLastPrice(db *influx.DB, query string) (float64, error) {
	result, err := db.GetData(query)
	if err != nil {
		return -1, fmt.Errorf("failed to get ohlc data from the influx db. Error %v", err)
	}

	lastPrice, err := parseLastPrice(result)
	if err != nil {
		return -1, fmt.Errorf("failed to parse the results. Error %v", err)

	}

	if len(*lastPrice) == 0 {
		return -1, fmt.Errorf("no last price data available. Error %v", err)
	}

	return (*lastPrice)[0].Price, nil
}

func (i Instrument) GetOHLC(db *influx.DB, query string) (*[]OHLC, error) {
	result, err := db.GetData(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get ohlc data from the influx db. Error %v", err)
	}

	ohlc, err := parseOHLC(result)
	if err != nil {
		return nil, fmt.Errorf("failed to parse the results. Error %v", err)

	}

	Logger.Info("OHCL result ", zap.Any("ohlc", ohlc))
	return ohlc, nil
}

func parseLastPrice(in *api.QueryTableResult) (*[]LastPrice, error) {
	out, err := parse(in)
	if err != nil {
		return nil, fmt.Errorf("failed to parse the results. Error %v", err)

	}
	lp := &[]LastPrice{}
	err = ms.Decode(out, lp)
	if err != nil {
		return nil, fmt.Errorf("failed to decode parsed db data into ohlc struct. Error %v", err)
	}

	return lp, nil
}

func parseOHLC(in *api.QueryTableResult) (*[]OHLC, error) {
	out, err := parse(in)
	if err != nil {
		return nil, fmt.Errorf("failed to parse the results. Error %v", err)

	}
	ohlc := &[]OHLC{}
	err = ms.Decode(out, ohlc)
	if err != nil {
		return nil, fmt.Errorf("failed to decode parsed db data into ohlc struct. Error %v", err)
	}

	return ohlc, nil
}
