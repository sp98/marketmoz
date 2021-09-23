package trade

import "github.com/sp98/marketmoz/pkg/strategy"

type Trade struct {
	Name         string
	NextPosition string
	Strategy     strategy.Strategy
}
