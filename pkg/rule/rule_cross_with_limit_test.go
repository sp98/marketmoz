package rule

import (
	"testing"

	"github.com/sp98/techan"
	"github.com/stretchr/testify/assert"
)

func TestCrossUpIndicatorRule(t *testing.T) {
	upInd := techan.NewFixedIndicator(3, 4, 5, 6)
	dnInd := techan.NewFixedIndicator(6, 5, 4, 3)

	rule := NewCrossUpWithLimitIndicatorRule(dnInd, upInd, 1)

	t.Run("always returns false when index == 0", func(t *testing.T) {
		assert.False(t, rule.IsSatisfied(0))
	})

	t.Run("Returns true when lower indicator crosses upper indicator with the max limit", func(t *testing.T) {
		assert.False(t, rule.IsSatisfied(1))
		assert.True(t, rule.IsSatisfied(2))
		assert.False(t, rule.IsSatisfied(3))
	})
}

func TestCrossDownIndicatorRule(t *testing.T) {
	upInd := techan.NewFixedIndicator(3, 4, 5, 6)
	dnInd := techan.NewFixedIndicator(6, 5, 4, 3)

	rule := NewCrossDownWithLimitIndicatorRule(dnInd, upInd, 1)

	t.Run("returns false when index == 0", func(t *testing.T) {
		assert.False(t, rule.IsSatisfied(0))
	})

	t.Run("returns true when upper indicator crosses below lower indicator", func(t *testing.T) {
		assert.False(t, rule.IsSatisfied(1))
		assert.True(t, rule.IsSatisfied(2))
		assert.False(t, rule.IsSatisfied(3))
	})
}
