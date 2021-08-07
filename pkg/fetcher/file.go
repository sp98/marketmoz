package fetcher

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go"
	"github.com/sp98/marketmoz/assets"
	"github.com/sp98/marketmoz/pkg/data"
	"github.com/sp98/marketmoz/pkg/db/influx"
	"github.com/sp98/marketmoz/pkg/utils"
	"go.uber.org/zap"
)

const (
	niftyAsset   = "data/nifty.txt"
	measurement  = "nifty-1m"
	organization = "marketmoz"
	bucket       = "test"
)

var tag = map[string]string{"stock": "nifty"}

const (
	INFLUXDB_URL   = "http://localhost:8086/"
	INFLUXDB_TOKEN = "m5txwvJXRbatNQM0AYKl9gkvtWVTkt_vIKU7IWotXQ-RAA-Q3i0wRrQfJTLvDmmn0e0GkCFJ0lZ3w8Pb-O_4uA=="
)

// startFileFetcher starts the fetching process from the file
func startFileFetcher() error {
	// Initialize Influx DB
	db := influx.NewDB(INFLUXDB_URL, INFLUXDB_TOKEN)
	writeAPI := db.Client.WriteAPIBlocking(organization, bucket)
	defer db.Client.Close()

	dataBytes, err := assets.ReadFile(niftyAsset)
	if err != nil {
		return fmt.Errorf("failed to read file %q", niftyAsset)
	}

	dataString := string(dataBytes)
	dataList := []data.FileData{}
	lines := strings.Split(dataString, "\n")
	for _, line := range lines {
		// Sleep for 1 second
		time.Sleep(1 * time.Second)
		l := strings.Split(line, ",")
		if len(l) > 6 {
			open, _ := strconv.ParseFloat(l[3], 64)
			high, _ := strconv.ParseFloat(l[4], 64)
			low, _ := strconv.ParseFloat(l[5], 64)
			close, _ := strconv.ParseFloat(l[6], 64)
			time, err := parseTime(l[1], l[2])
			if err != nil {
				Logger.Error("failed to parse time", zap.Error(err))
				return err
			}

			d := data.FileData{
				Stock: l[0],
				Open:  open,
				High:  high,
				Low:   low,
				Close: close,
				Time:  time,
			}

			Logger.Info("Tick", zap.Object("tick data", &d))

			fields := map[string]interface{}{
				"Open":  d.Open,
				"High":  d.High,
				"Low":   d.Low,
				"Close": d.Close,
			}

			p := influxdb2.NewPoint(measurement, tag, fields, d.Time)
			err = writeAPI.WritePoint(context.Background(), p)
			if err != nil {
				Logger.Error("failed to write point", zap.Object("point", &d), zap.Error(err))
			}
			dataList = append(dataList, d)
		}
	}

	Logger.Info("Data List created", zap.Int("count", len(dataList)))

	return nil
}

func parseTime(d, t string) (time.Time, error) {
	// Time format in the text file is 20210101,09:08
	// We have to convert it to 20210101090800
	parsedtime, err := utils.ToTime(fmt.Sprintf("%s%s%s", d, strings.ReplaceAll(t, ":", ""), "00"))
	return parsedtime, err
}
