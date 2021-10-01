package strategy

import (
	"fmt"
	"time"

	"github.com/sdcoffey/big"
	"github.com/sp98/marketmoz/pkg/rule"
	"github.com/sp98/marketmoz/pkg/utils"
	"github.com/sp98/techan"
)

// BasicEma is an example of how to create a basic Exponential moving average indicator
// based on the close prices of a timeseries from your exchange of choice.

const (
	EMA21 = 21
	EMA10 = 10
	RSI14 = 14
)

// func GetSeries() (*techan.TimeSeries, int) {

// 	series := techan.NewTimeSeries()

// 	data := data.NewOHLCData("marketmoz", "2112121 ", "1m", "NSE", "EQ")
// 	ohlc, _ := data.GetOHLC()
// 	d := ohlc.GetData()

// 	count := 0
// 	for _, datum := range *d {
// 		period := techan.NewTimePeriod(time.Unix(datum.Time, 0), time.Minute*1)
// 		candle := techan.NewCandle(period)
// 		candle.OpenPrice = big.NewFromString(fmt.Sprintf("%f", datum.Open))
// 		candle.ClosePrice = big.NewFromString(fmt.Sprintf("%f", datum.Close))
// 		candle.MaxPrice = big.NewFromString(fmt.Sprintf("%f", datum.High))
// 		candle.MinPrice = big.NewFromString(fmt.Sprintf("%f", datum.Low))
// 		series.AddCandle(candle)
// 		count = count + 1
// 	}

// 	return series, count

// }
func BasicEma(window int, series *techan.TimeSeries) techan.Indicator {
	closePrices := techan.NewClosePriceIndicator(series)
	movingAverage := techan.NewEMAIndicator(closePrices, window)
	return movingAverage
}

func BasicRSI(timeframe int, series *techan.TimeSeries) techan.Indicator {
	closePrices := techan.NewClosePriceIndicator(series)
	rsi := techan.NewRelativeStrengthIndexIndicator(closePrices, timeframe)
	return rsi
}

func getSeriesFromCSV(records [][]string) (*techan.TimeSeries, error) {
	series := techan.NewTimeSeries()
	for _, record := range records {
		date, err := time.Parse("2006-01-02", record[0])
		if err != nil {
			return nil, err
		}
		period := techan.NewTimePeriod(date, time.Hour*24)
		candle := techan.NewCandle(period)
		candle.OpenPrice = big.NewFromString(record[1])
		candle.MaxPrice = big.NewFromString(record[2])
		candle.MinPrice = big.NewFromString(record[3])
		candle.ClosePrice = big.NewFromString(record[4])
		candle.Volume = big.NewFromString(record[6])

		series.AddCandle(candle)

	}

	return series, nil
}

func ExampleStrategy() {
	records, err := utils.CSVReader("./assets/data/RELIANCE.NS.csv")
	if err != nil {
		fmt.Println("failed to read records. Error : ", err)
		return
	}

	series, err := getSeriesFromCSV(records)
	if err != nil {
		fmt.Println("failed to get series. Error:: ", err)
		return
	}

	closePrices := techan.NewClosePriceIndicator(series)
	volume := techan.NewVolumeIndicator(series)

	pvtIndicator := techan.NewPriceVolumeTrendIndicator(closePrices, volume, 1)

	// subtract 2 because first row in the csv is header.
	fmt.Println("PVT - ", pvtIndicator.Calculate(len(records)-1))

	pvtSignalIndicator := techan.NewPVTAndSignalIndicator(pvtIndicator, techan.NewEMAIndicator(pvtIndicator, 21))
	fmt.Println("PVT Signal Diff- ", pvtSignalIndicator.Calculate(len(records)-1))

	//RSI
	rsiIndicator := techan.NewRelativeStrengthIndexIndicator(closePrices, 14)
	fmt.Println("RSI - ", rsiIndicator.Calculate(len(records)-1))

	macdIndicator := techan.NewMACDIndicator(closePrices, 12, 26)
	fmt.Println("MACD - ", macdIndicator.Calculate(len(records)-1))

	macdHistogramIndicator := techan.NewMACDHistogramIndicator(techan.NewMACDIndicator(closePrices, 12, 26), 9)
	fmt.Println("MACD Histogram - ", macdHistogramIndicator.Calculate(len(records)-1))

	stochasticFastIndicator := techan.NewFastStochasticIndicator(series, 14)
	fmt.Println("fast Stochastic Indicator", stochasticFastIndicator.Calculate(len(records)-1))

	stochasticSlowIndicator := techan.NewSlowStochasticIndicator(stochasticFastIndicator, 3)
	fmt.Println("slow Stochastic Indicator", stochasticSlowIndicator.Calculate(len(records)-1))

	pivotPoint := techan.NewPivoPointIndicator(series, techan.DAY)
	fmt.Println("pivot point indicator", pivotPoint.Calculate(len(records)-1))

	pivotLevelR1 := techan.NewPivotLevelIndicator(series, techan.DAY, techan.RESISTANCE_1)
	fmt.Println("pivot level R1", pivotLevelR1.Calculate(len(records)-1))

	pivotLevelR2 := techan.NewPivotLevelIndicator(series, techan.DAY, techan.RESISTANCE_2)
	fmt.Println("pivot level R2", pivotLevelR2.Calculate(len(records)-1))

	pivotLevelR3 := techan.NewPivotLevelIndicator(series, techan.DAY, techan.RESISTANCE_3)
	fmt.Println("pivot level R3", pivotLevelR3.Calculate(len(records)-1))

	pivotLevelS1 := techan.NewPivotLevelIndicator(series, techan.DAY, techan.SUPPORT_1)
	fmt.Println("pivot level S1", pivotLevelS1.Calculate(len(records)-1))

	pivotLevelS2 := techan.NewPivotLevelIndicator(series, techan.DAY, techan.SUPPORT_2)
	fmt.Println("pivot level S2", pivotLevelS2.Calculate(len(records)-1))

	pivotLevelS3 := techan.NewPivotLevelIndicator(series, techan.DAY, techan.SUPPORT_3)
	fmt.Println("pivot level S3", pivotLevelS3.Calculate(len(records)-1))

	closePrice := techan.NewClosePriceIndicator(series).Calculate(len(records) - 1)
	fmt.Println("last close price: ", closePrice)
	buyPivotPointRule := rule.NewPivotPointRule(series, closePrice, techan.MONTH, "BUY").IsSatisfied(len(records) - 1)
	sellPivotPointRule := rule.NewPivotPointRule(series, closePrice, techan.MONTH, "SELL").IsSatisfied(len(records) - 1)

	fmt.Println("Buy based on pivot point rule: ", buyPivotPointRule)
	fmt.Println("Sell based on pivot point rule: ", sellPivotPointRule)

	s1 := PVTStrategyRules(series)
	fmt.Printf("Strategy - %+v\n", s1)
	fmt.Println("Should Enter Long: ", s1.ShouldEnterLong(len(records)-1))
	fmt.Println("Should Enter Short: ", s1.ShouldEnterShort(len(records)-1))

	trend := techan.NewTrendlineIndicator(closePrices, 5)
	fmt.Println("Trend Line: ", trend.Calculate(len(records)-1))

	isBullishMarubozu := techan.NewBullishMarubozuIndicator(series, big.NewFromInt(90))
	fmt.Println("is Bullish Marubozu: ", isBullishMarubozu.Calculate(len(records)-1))

	isBearishMarubozu := techan.NewBearishMarubozuIndicator(series, big.NewFromInt(90))
	fmt.Println("is bearish Marubozu: ", isBearishMarubozu.Calculate(len(records)-1))
}
