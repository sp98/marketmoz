package trade

import (
	"fmt"
	"time"

	"github.com/sdcoffey/big"
	"github.com/sp98/techan"
)

type Series struct {
	next Flow
}

func (s *Series) Execute(t *Trade) {
	fmt.Println("Flow: Get Time Series data")
	// Get OHLC data and build the strategy rules
	series, err := getSeries(t)
	if err != nil {
		fmt.Println("failed to get ohlc data for the instrument - ", t.Instrument.Name)
		return
	}
	// TODO: create cache of the series and only add new candle every minute. If cache is empty then get new series everytime
	t.Series = series

	if s.next != nil {
		s.next.Execute(t)
	}
}

func (s *Series) SetNext(next Flow) {
	s.next = next
}

func getSeries(t *Trade) (*techan.TimeSeries, error) {
	series := techan.NewTimeSeries()
	ohlcList, err := t.Instrument.GetOHLC(t.DB)
	if err != nil {
		return nil, err
	}
	for _, datum := range *ohlcList {
		period := techan.NewTimePeriod(time.Unix(datum.Time, 0), time.Minute*1)
		candle := techan.NewCandle(period)
		candle.OpenPrice = big.NewFromString(fmt.Sprintf("%f", datum.Open))
		candle.ClosePrice = big.NewFromString(fmt.Sprintf("%f", datum.Close))
		candle.MaxPrice = big.NewFromString(fmt.Sprintf("%f", datum.High))
		candle.MinPrice = big.NewFromString(fmt.Sprintf("%f", datum.Low))
		series.AddCandle(candle)
	}

	return series, nil
}
