package common

type TokenDetails struct {
	Exchange       string
	InstrumentType string
	Segment        string
}

var tokenMap = map[string]TokenDetails{
	"738561": {
		Exchange:       "NSE",
		InstrumentType: "NSE",
		Segment:        "EQ",
	},

	"128083204": {
		Exchange:       "BSE",
		InstrumentType: "BSE",
		Segment:        "EQ",
	},
}

func GetTokenMap() *map[string]TokenDetails {
	return &tokenMap
}

func GetTokenDetails(token string) *TokenDetails {
	td, ok := tokenMap[token]
	if !ok {
		return nil
	}
	return &td
}
