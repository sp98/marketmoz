package strategy

import (
	"github.com/sp98/marketmoz/pkg/rule"
)

type Strategy interface {
	ShouldEnterLong(index int) bool
	ShouldEnterShort(index int) bool
	ShouldExitLong(index int) bool
	ShouldExistShort(index int) bool
}

type RuleStrategy struct {
	LongEntryRule  rule.Rule
	LongExitRule   rule.Rule
	ShortEntryRule rule.Rule
	ShortExitRule  rule.Rule
}

func (rs RuleStrategy) ShouldEnterLong(index int) bool {
	if rs.LongEntryRule == nil {
		panic("Long Entry rule can't be nil")
	}

	return rs.LongEntryRule.IsSatisfied(index)
}

func (rs RuleStrategy) ShouldExitLong(index int) bool {
	if rs.LongExitRule == nil {
		panic("Long Exit rule can't be nil")
	}

	return rs.LongExitRule.IsSatisfied(index)
}

func (rs RuleStrategy) ShouldEnterShort(index int) bool {
	if rs.ShortEntryRule == nil {
		panic("Short Entry rule can't be nil")
	}

	return rs.ShortEntryRule.IsSatisfied(index)
}

func (rs RuleStrategy) ShouldExistShort(index int) bool {
	if rs.ShortExitRule == nil {
		panic("Short Exit rule can't be nil")
	}

	return rs.ShortExitRule.IsSatisfied(index)
}
