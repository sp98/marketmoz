package common

type Instrument struct {
	Name           string
	Symbol         string
	Token          string
	Exchange       string
	InstrumentType string
	Segment        string
}

var InstrumentMap = map[string]Instrument{
	"61348": {
		Name:           "RELIANCE INDUSTRIES",
		Symbol:         "RELIANCE",
		Token:          "61348",
		Exchange:       "NSE",
		InstrumentType: "EQ",
		Segment:        "NSE",
	},
	"59393": {
		Name:           "HAVELLS INDIA",
		Symbol:         "HAVELLS",
		Token:          "59393",
		Exchange:       "NSE",
		InstrumentType: "EQ",
		Segment:        "NSE",
	},
	"59163": {
		Name:           "INFOSYS",
		Symbol:         "INFY",
		Token:          "59163",
		Exchange:       "NSE",
		InstrumentType: "EQ",
		Segment:        "NSE",
	},

	//TODO: Add more instruments
}

func GetInstrumentMap() *map[string]Instrument {
	return &InstrumentMap
}

func GetInstrumentDetails(token string) *Instrument {
	instrument, ok := InstrumentMap[token]
	if !ok {
		return nil
	}
	return &instrument
}
