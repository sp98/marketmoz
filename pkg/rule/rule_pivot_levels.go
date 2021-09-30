package rule

import (
	"github.com/sdcoffey/big"
	"github.com/sp98/techan"
)

type pivotPointRule struct {
	lastPrice  big.Decimal
	pivotPoint techan.Indicator
	r1         techan.Indicator
	r2         techan.Indicator
	r3         techan.Indicator
	s1         techan.Indicator
	s2         techan.Indicator
	s3         techan.Indicator
	trade      string
}

func NewPivotPointRule(series *techan.TimeSeries, lastPrice big.Decimal, timeLevel techan.TimeLevel, trade string) Rule {
	return pivotPointRule{
		lastPrice:  lastPrice,
		pivotPoint: techan.NewPivoPointIndicator(series, timeLevel),
		r1:         techan.NewPivotLevelIndicator(series, timeLevel, techan.RESISTANCE_1),
		r2:         techan.NewPivotLevelIndicator(series, timeLevel, techan.RESISTANCE_2),
		r3:         techan.NewPivotLevelIndicator(series, timeLevel, techan.RESISTANCE_3),
		s1:         techan.NewPivotLevelIndicator(series, timeLevel, techan.SUPPORT_1),
		s2:         techan.NewPivotLevelIndicator(series, timeLevel, techan.SUPPORT_2),
		s3:         techan.NewPivotLevelIndicator(series, timeLevel, techan.SUPPORT_3),
		trade:      trade,
	}
}

func (r pivotPointRule) IsSatisfied(index int) bool {
	r1 := r.r1.Calculate(index)
	r2 := r.r2.Calculate(index)
	r3 := r.r3.Calculate(index)
	p := r.pivotPoint.Calculate(index)
	s1 := r.s1.Calculate(index)
	s2 := r.s2.Calculate(index)
	s3 := r.s3.Calculate(index)

	var upper big.Decimal
	var lower big.Decimal
	switch {
	case r.lastPrice.LTE(s3):
	case r.lastPrice.GTE(s3) && r.lastPrice.LT(s2):
		upper = s2
		lower = s3
	case r.lastPrice.GTE(s2) && r.lastPrice.LT(s1):
		upper = s1
		lower = s2
	case r.lastPrice.GTE(s1) && r.lastPrice.LT(p):
		upper = p
		lower = s1
	case r.lastPrice.GTE(p) && r.lastPrice.LT(r1):
		upper = r1
		lower = p
	case r.lastPrice.GTE(r1) && r.lastPrice.LT(r2):
		upper = r2
		lower = r1
	case r.lastPrice.GTE(r2) && r.lastPrice.LT(r3):
		upper = r3
		lower = r2
	case r.lastPrice.GTE(r3):
	}

	percentLevel := big.NewFromInt(100).Mul(lower.Sub(r.lastPrice)).Div(upper.Sub(lower))

	switch r.trade {
	case "BUY":
		return percentLevel.LTE(big.NewFromInt(75)) && percentLevel.GTE(big.NewFromInt(5))
	case "SELL":
		return percentLevel.GTE(big.NewFromInt(25)) && percentLevel.LTE(big.NewFromInt(95))
	}

	return false
}
