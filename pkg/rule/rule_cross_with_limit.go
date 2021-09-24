package rule

import "github.com/sp98/techan"

// NewCrossUpWithLimitIndicatorRule returns a new rule that is satisfied when the lower indicator has crossed above the upper
// indicator within the provide limit. Limit should be greater than 0.
func NewCrossUpWithLimitIndicatorRule(upper, lower techan.Indicator, limit int) Rule {
	return crossWithLimitRule{
		upper:    upper,
		lower:    lower,
		cmp:      1,
		maxLimit: limit,
	}
}

// NewCrossDownWithLimitIndicatorRule returns a new rule that is satisfied when the upper indicator has crossed below the lower
// indicator within the provided limit. Limit should be greater than 0.
func NewCrossDownWithLimitIndicatorRule(upper, lower techan.Indicator, limit int) Rule {
	return crossWithLimitRule{
		upper:    lower,
		lower:    upper,
		cmp:      -1,
		maxLimit: limit,
	}
}

type crossWithLimitRule struct {
	upper    techan.Indicator
	lower    techan.Indicator
	cmp      int
	maxLimit int
}

func (clr crossWithLimitRule) IsSatisfied(index int) bool {
	i := index

	if i == 0 {
		return false
	}

	totalSatisfied := 0
	if cmp := clr.lower.Calculate(i).Cmp(clr.upper.Calculate(i)); cmp == 0 || cmp == clr.cmp {
		for ; i >= 0; i-- {
			if totalSatisfied > clr.maxLimit {
				return false
			}
			if cmp = clr.lower.Calculate(i).Cmp(clr.upper.Calculate(i)); cmp == 0 || cmp == -clr.cmp {
				return true
			}
			totalSatisfied++
		}
	}

	return false
}
