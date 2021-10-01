package data

var InstrumentMap = map[string]Instrument{
	"408065": {
		Name:           "INFOSYS",
		Symbol:         "INFY",
		Token:          uint32(408065),
		Exchange:       "NSE",
		InstrumentType: "EQ",
		Segment:        "NSE",
	},
	"93534": {
		Name:           "RELIANCE INDUSTRIES",
		Symbol:         "RELIANCE",
		Token:          uint32(93534),
		Exchange:       "NSE",
		InstrumentType: "EQ",
		Segment:        "NSE",
	},
	"59393": {
		Name:           "HAVELLS INDIA",
		Symbol:         "HAVELLS",
		Token:          uint32(59393),
		Exchange:       "NSE",
		InstrumentType: "EQ",
		Segment:        "NSE",
	},
	"59163": {
		Name:           "INFOSYS",
		Symbol:         "INFY",
		Token:          uint32(59163),
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
