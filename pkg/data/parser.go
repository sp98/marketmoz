package data

import (
	"fmt"

	"github.com/influxdata/influxdb-client-go/v2/api"
)

func parse(in *api.QueryTableResult) ([]map[string]interface{}, error) {
	out := []map[string]interface{}{}
	currentField := ""
	index := 0
	for in.Next() {
		if currentField != in.Record().Field() {
			currentField = in.Record().Field()
			index = 0
		}

		if len(out) < index+1 {
			r := map[string]interface{}{}
			r[in.Record().Field()] = in.Record().Value()
			out = append(out, r)
		} else {
			out[index][currentField] = in.Record().Value()
			out[index]["Time"] = in.Record().Time().Local().Unix()
		}

		index += 1
	}
	// check for an error
	if in.Err() != nil {
		fmt.Printf("query parsing error: %s\n", in.Err().Error())
		return nil, in.Err()
	}

	return out, nil
}
