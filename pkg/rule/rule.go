package rule

type Rule interface {
	IsSatisfied(index int) bool
}

// AndRule is rule that is satisfied only when all the passed in rules are satisfied
type AndRule struct {
	r []Rule
}

func NewAndRule(r ...Rule) Rule {
	return AndRule{r}
}

func (ar AndRule) IsSatisfied(index int) bool {
	for _, r := range ar.r {
		if !r.IsSatisfied(index) {
			return false
		}
	}
	return true
}

// OrRule is a rule that is satisfied when any of the passed in rule is satisfied.
type OrRule struct {
	r []Rule
}

func NewOrRule(r ...Rule) Rule {
	return OrRule{r}
}

func (or OrRule) IsSatisfied(index int) bool {
	for _, r := range or.r {
		if r.IsSatisfied(index) {
			return true
		}
	}
	return false
}

type AndOrRule struct {
	andRules []Rule
	orRules  []Rule
}

func (aor *AndOrRule) SetAndRule(ar ...Rule) {
	aor.andRules = ar
}

func (aor *AndOrRule) SetOrRule(or ...Rule) {
	aor.orRules = or
}

func (aor *AndOrRule) IsSatisfied(index int) bool {
	// TODO: evaluate in a goroutine
	andRulesSatisfied := true
	for _, r := range aor.andRules {
		if !r.IsSatisfied(index) {
			// No need to evaluate futher if any of the AndRules is false
			return false
		}
	}

	orRulesSatisfied := false
	for _, r := range aor.orRules {
		if r.IsSatisfied(index) {
			orRulesSatisfied = true
			// No need to evaluate further rules since one of the orRules is satisfied
			break
		}
	}

	return andRulesSatisfied && orRulesSatisfied
}
