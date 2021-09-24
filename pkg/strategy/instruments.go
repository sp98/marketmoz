package strategy

type Instrument struct {
	Name           string
	Symbol         string
	Token          string
	Exchange       string
	InstrumentType string
	Segment        string
}

type Instruments []Instrument

// InstrumentStrategyMap specifies the instruments to be traded for a particular stragegy
var InstrumentStrategyMap = map[string]Instruments{
	"PVT": {
		// TODO: Add all the Instruments to be traded for the PVT Strategy
		{
			Name:     "RELIANCE INDUSTRIES",
			Symbol:   "RELIANCE",
			Token:    "61348",
			Exchange: "NSE",
		},
	},
}
