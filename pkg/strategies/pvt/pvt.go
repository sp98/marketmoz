package pvt

import (
	"fmt"
	"time"

	"github.com/sdcoffey/big"
	"github.com/sp98/marketmoz/pkg/data"
	"github.com/sp98/techan"
)

// BasicEma is an example of how to create a basic Exponential moving average indicator
// based on the close prices of a timeseries from your exchange of choice.

const (
	EMA21 = 21
	EMA10 = 10
	RSI14 = 14
)

func GetSeries() (*techan.TimeSeries, int) {

	series := techan.NewTimeSeries()

	data := data.NewOHLCData("marketmoz", "2112121 ", "1m", "NSE", "EQ")
	ohlc, _ := data.GetOHLC()
	d := ohlc.GetData()

	count := 0
	for _, datum := range *d {
		period := techan.NewTimePeriod(time.Unix(datum.Time, 0), time.Minute*1)
		candle := techan.NewCandle(period)
		candle.OpenPrice = big.NewFromString(fmt.Sprintf("%f", datum.Open))
		candle.ClosePrice = big.NewFromString(fmt.Sprintf("%f", datum.Close))
		candle.MaxPrice = big.NewFromString(fmt.Sprintf("%f", datum.High))
		candle.MinPrice = big.NewFromString(fmt.Sprintf("%f", datum.Low))
		series.AddCandle(candle)
		count = count + 1
		fmt.Println("time - ", time.Unix(datum.Time, 0))
	}

	return series, count

}
func BasicEma(window int, series *techan.TimeSeries) techan.Indicator {

	closePrices := techan.NewClosePriceIndicator(series)
	fmt.Println("series - ", series)
	movingAverage := techan.NewEMAIndicator(closePrices, window) // Create an exponential moving average with a window of 10

	//sma := techan.NewSimpleMovingAverage(closePrices, 1)
	return movingAverage
}

func BasicRSI(timeframe int, series *techan.TimeSeries) techan.Indicator {
	closePrices := techan.NewClosePriceIndicator(series)
	fmt.Println("series - ", series)
	rsi := techan.NewRelativeStrengthIndexIndicator(closePrices, timeframe)
	return rsi
}

func StrategyExample1() {
	series, index := GetSeries()
	if index < EMA21 {
		fmt.Printf("not enough data to creat a 21 period EMA. Required %d. Present %d", EMA21, index)
		return
	}
	indicator := BasicEma(EMA21, series)
	// record trades on this object
	record := techan.NewTradingRecord()

	entryConstant := techan.NewConstantIndicator(14020)
	exitConstant := techan.NewConstantIndicator(10)

	entryRule := techan.And(
		techan.NewCrossUpIndicatorRule(entryConstant, indicator),
		techan.PositionNewRule{}) // Is satisfied when the price ema moves above 30 and the current position is new

	exitRule := techan.And(
		techan.NewCrossDownIndicatorRule(indicator, exitConstant),
		techan.PositionOpenRule{}) // Is satisfied when the price ema moves below 10 and the current position is open

	strategy := techan.RuleStrategy{
		UnstablePeriod: 5,
		EntryRule:      entryRule,
		ExitRule:       exitRule,
	}

	fmt.Println(strategy.ShouldEnter(index-1, record)) // returns false
}

func StrategyExample2() {
	series, index := GetSeries()
	if index < EMA21 {
		fmt.Printf("not enough data to creat a 21 period EMA. Required %d. Present %d", EMA21, index)
		return
	}
	indicator1 := BasicEma(EMA21, series)
	indicator2 := BasicEma(EMA10, series)
	// record trades on this object
	record := techan.NewTradingRecord()

	// entryConstant := techan.NewConstantIndicator(14020)
	// exitConstant := techan.NewConstantIndicator(10)

	entryRule := techan.And(
		techan.NewCrossUpIndicatorRule(indicator2, indicator1),
		techan.PositionNewRule{}) // Is satisfied when the price ema moves above 30 and the current position is new

	exitRule := techan.And(
		techan.NewCrossDownIndicatorRule(indicator1, indicator2),
		techan.PositionOpenRule{}) // Is satisfied when the price ema moves below 10 and the current position is open

	strategy := techan.RuleStrategy{
		UnstablePeriod: 5,
		EntryRule:      entryRule,
		ExitRule:       exitRule,
	}

	fmt.Println(strategy.ShouldEnter(index-1, record)) // returns false
}

func StrategyExample3() {
	series, index := GetSeries()
	if index < RSI14 {
		fmt.Printf("not enough data to create a 14 period RSI. Required %d. Present %d", RSI14, index)
		return
	}
	indicator1 := BasicRSI(RSI14, series)
	// record trades on this object
	record := techan.NewTradingRecord()

	entryConstant := techan.NewConstantIndicator(50)
	exitConstant := techan.NewConstantIndicator(50)

	entryRule := techan.And(
		techan.NewCrossUpIndicatorRule(entryConstant, indicator1),
		techan.PositionNewRule{}) // Is satisfied when the price ema moves above 30 and the current position is new

	exitRule := techan.And(
		techan.NewCrossDownIndicatorRule(indicator1, exitConstant),
		techan.PositionOpenRule{}) // Is satisfied when the price ema moves below 10 and the current position is open

	strategy := techan.RuleStrategy{
		UnstablePeriod: 5,
		EntryRule:      entryRule,
		ExitRule:       exitRule,
	}

	fmt.Println(strategy.ShouldEnter(index-1, record)) // returns false
}
