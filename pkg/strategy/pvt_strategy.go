package strategy

import (
	"github.com/sp98/marketmoz/pkg/data"
	"github.com/sp98/marketmoz/pkg/rule"
	"github.com/sp98/techan"
)

var PVTInstruments = []string{"61348", "59393", "59163"}

func PVTStrategyRules(series *techan.TimeSeries) Strategy {
	// Create RuleStrategy using all entry and exit rules for PVT strategy
	closePriceIndicator := techan.NewClosePriceIndicator(series)
	volumeIndicator := techan.NewVolumeIndicator(series)

	pvtIndicator := techan.NewPriceVolumeTrendIndicator(closePriceIndicator, volumeIndicator, 1)
	pvtEMAIndicator := techan.NewEMAIndicator(pvtIndicator, 21)

	rsiIndicator := techan.NewRelativeStrengthIndexIndicator(closePriceIndicator, 14)
	rsiConstantIndicator := techan.NewConstantIndicator(50)

	macdHistogramIndicator := techan.NewMACDHistogramIndicator(techan.NewMACDIndicator(closePriceIndicator, 12, 26), 9)
	macdHistogramConstantIndicator := techan.NewConstantIndicator(0)
	longEntryRule := rule.NewAndRule(
		rule.NewCrossUpWithLimitIndicatorRule(pvtEMAIndicator, pvtIndicator, 1),
		rule.NewCrossUpWithLimitIndicatorRule(rsiConstantIndicator, rsiIndicator, 5),
		rule.NewCrossUpWithLimitIndicatorRule(macdHistogramConstantIndicator, macdHistogramIndicator, 5),
	)

	shortEntryRule := rule.NewAndRule(
		rule.NewCrossUpWithLimitIndicatorRule(pvtIndicator, pvtEMAIndicator, 1),
		rule.NewCrossUpWithLimitIndicatorRule(rsiIndicator, rsiConstantIndicator, 5),
		rule.NewCrossUpWithLimitIndicatorRule(macdHistogramIndicator, macdHistogramConstantIndicator, 5),
	)

	return &RuleStrategy{
		LongEntryRule:  longEntryRule,
		ShortEntryRule: shortEntryRule,
	}
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
