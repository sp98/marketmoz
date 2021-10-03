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
	"738561": {
		Name:           "RELIANCE INDUSTRIES",
		Symbol:         "RELIANCE",
		Token:          uint32(738561),
		Exchange:       "NSE",
		InstrumentType: "EQ",
		Segment:        "NSE",
	},
	"5633": {
		Name:           "ACC",
		Symbol:         "ACC",
		Token:          uint32(5633),
		Exchange:       "NSE",
		InstrumentType: "EQ",
		Segment:        "NSE",
	},
	"2513665": {
		Name:           "HAVELLS INDIA",
		Symbol:         "HAVELLS",
		Token:          uint32(2513665),
		Exchange:       "NSE",
		InstrumentType: "EQ",
		Segment:        "NSE",
	},
	"492033": {
		Name:           "KOTAK MAHINDRA BANK",
		Symbol:         "KOTAKBANK",
		Token:          uint32(492033),
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
