package strategy

import (
	"github.com/sp98/marketmoz/pkg/data"
	"github.com/sp98/marketmoz/pkg/rule"
	"github.com/sp98/techan"
)

var PVTInstruments = []string{"61348", "59393", "59163"}

func PVTStrategyRules(series *techan.TimeSeries) Strategy {

	// Set indicators
	closePriceIndicator := techan.NewClosePriceIndicator(series)
	volumeIndicator := techan.NewVolumeIndicator(series)

	pvtIndicator := techan.NewPriceVolumeTrendIndicator(closePriceIndicator, volumeIndicator, 1)
	pvtEMAIndicator := techan.NewEMAIndicator(pvtIndicator, 21)

	rsiIndicator := techan.NewRelativeStrengthIndexIndicator(closePriceIndicator, 14)
	rsiConstantIndicator := techan.NewConstantIndicator(50)

	macdHistogramIndicator := techan.NewMACDHistogramIndicator(techan.NewMACDIndicator(closePriceIndicator, 12, 26), 9)
	macdHistogramConstantIndicator := techan.NewConstantIndicator(0)

	isBullish := techan.NewBullishIndicator(series)
	isBearish := techan.NewBearishIndicator(series)

	// Set rules

	// Set Long Entry rule
	longEntryRule := &rule.AndOrRule{}
	longEntryRule.SetAndRule(
		rule.NewCrossUpWithLimitIndicatorRule(pvtEMAIndicator, pvtIndicator, 1),
		rule.NewCrossUpWithLimitIndicatorRule(rsiConstantIndicator, rsiIndicator, 2),
		rule.NewCrossUpWithLimitIndicatorRule(macdHistogramConstantIndicator, macdHistogramIndicator, 5),
	)

	longEntryRule.SetOrRule(
		rule.NewEqualRule(isBullish, techan.NewConstantIndicator(1)),
	)

	// Set Short Entry rule
	shortEntryRule := &rule.AndOrRule{}
	shortEntryRule.SetAndRule(
		rule.NewCrossDownWithLimitIndicatorRule(pvtIndicator, pvtEMAIndicator, 1),
		rule.NewCrossDownWithLimitIndicatorRule(rsiIndicator, rsiConstantIndicator, 2),
		rule.NewCrossDownWithLimitIndicatorRule(macdHistogramIndicator, macdHistogramConstantIndicator, 5),
	)

	shortEntryRule.SetOrRule(
		rule.NewEqualRule(isBearish, techan.NewConstantIndicator(1)),
	)

	return &RuleStrategy{
		LongEntryRule:  longEntryRule,
		ShortEntryRule: shortEntryRule,
	}
}

// GetPVTStrategyLongSL returns the lowest between current candle low and previous candle low
func GetPVTStrategyLongSL(series *techan.TimeSeries) float64 {
	index := series.LastIndex()
	currentLow := techan.NewLowPriceIndicator(series).Calculate(index)
	previousLow := techan.NewLowPriceIndicator(series).Calculate(index - 1)

	if currentLow.LT(previousLow) {
		return currentLow.Float()
	}
	return previousLow.Float()

}

// GetPVTStrategyShortSL returns the highest between current candle high and previous candle high
func GetPVTStrategyShortSL(series *techan.TimeSeries) float64 {
	index := series.LastIndex()
	currentHigh := techan.NewHighPriceIndicator(series).Calculate(index)
	previousHigh := techan.NewHighPriceIndicator(series).Calculate(index - 1)

	if currentHigh.GT(previousHigh) {
		return currentHigh.Float()
	}
	return previousHigh.Float()
}

//GetPVTInstruments returns a list of instruments that should be traded with PVT strategy
func GetPVTInstruments() *[]data.Instrument {
	pvtInstruments := []data.Instrument{}

	for _, token := range PVTInstruments {
		instrument := data.GetInstrumentDetails(token)
		if instrument != nil {
			pvtInstruments = append(pvtInstruments, *instrument)
		}
	}

	return &pvtInstruments
}
