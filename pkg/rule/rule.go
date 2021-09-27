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
	andRulesSatisfied := false
	for _, r := range aor.andRules {
		if r.IsSatisfied(index) {
			andRulesSatisfied = true
		}
	}

	orRulesSatisfied := false
	for _, r := range aor.orRules {
		if r.IsSatisfied(index) {
			orRulesSatisfied = true
		}
	}

	return andRulesSatisfied && orRulesSatisfied
}
