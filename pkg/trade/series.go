package trade

import (
	"fmt"
	"time"

	"github.com/sdcoffey/big"
	"github.com/sp98/marketmoz/assets"
	"github.com/sp98/marketmoz/pkg/common"
	"github.com/sp98/techan"
	"go.uber.org/zap"
)

type Series struct {
	next Flow
}

func (s *Series) Execute(t *Trade) error {
	Logger.Info("Flow: Get Time Series data")
	// Get OHLC data and build the strategy rules
	series, err := getSeries(t)
	if err != nil {
		return fmt.Errorf("failed to get ohlc data for the instrument %q. Error %v", t.Instrument.Name, err)
	}
	// TODO: create cache of the series and only add new candle every minute. If cache is empty then get new series everytime
	t.Series = series

	if s.next != nil {
		return s.next.Execute(t)
	}

	return nil
}

func (s *Series) SetNext(next Flow) {
	s.next = next
}

func getSeries(t *Trade) (*techan.TimeSeries, error) {
	series := techan.NewTimeSeries()
	query, err := t.Instrument.GetQuery(t.Interval, common.OHLC_QUERY_ASSET)
	//query, err := GetTestQuery()
	if err != nil {
		Logger.Error("failed to get query", zap.Error(err))
		return nil, err
	}

	ohlcList, err := t.Instrument.GetOHLC(t.DB, query)
	if err != nil {
		return nil, err
	}
	interval := t.GetIntervalTime()
	for _, datum := range *ohlcList {
		period := techan.NewTimePeriod(time.Unix(datum.Time, 0), interval)
		candle := techan.NewCandle(period)
		candle.OpenPrice = big.NewFromString(fmt.Sprintf("%f", datum.Open))
		candle.ClosePrice = big.NewFromString(fmt.Sprintf("%f", datum.Close))
		candle.MaxPrice = big.NewFromString(fmt.Sprintf("%f", datum.High))
		candle.MinPrice = big.NewFromString(fmt.Sprintf("%f", datum.Low))
		series.AddCandle(candle)
	}

	return series, nil
}

func GetTestQuery() (string, error) {
	queryBytes, err := assets.ReadFileAndReplace(
		common.OHLC_QUERY_TEST_ASSET,
		[]string{},
	)
	return string(queryBytes), err
}
