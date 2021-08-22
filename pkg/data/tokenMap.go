package data

import (
	"fmt"

	"github.com/sp98/marketmoz/pkg/common"
)

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

func GetBucketName(token string) string {
	tm, ok := tokenMap[token]
	if !ok {
		return ""
	}

	b := fmt.Sprintf(common.REAL_TIME_DATA_BUCKET, tm.Exchange, tm.Segment)
	return b
}
