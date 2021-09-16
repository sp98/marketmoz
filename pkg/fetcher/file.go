package fetcher

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/influxdata/influxdb-client-go/v2/domain"
	"github.com/sp98/marketmoz/assets"
	"github.com/sp98/marketmoz/pkg/common"
	"github.com/sp98/marketmoz/pkg/db/influx"
	"github.com/sp98/marketmoz/pkg/utils"
	"go.uber.org/zap"
)

const (
	niftyAsset  = "data/nifty.txt"
	measurement = "nifty-lp"
	bucket      = "test"
)

var tag = map[string]string{"stock": "nifty"}

// startFileFetcher starts the fetching process from the file
func startFileFetcher() error {

	ctx := context.Background()
	// Initialize Influx DB
	db := influx.NewDB(ctx, common.INFLUXDB_ORGANIZATION, common.INFLUXDB_URL, common.INFLUXDB_TOKEN)
	defer db.Client.Close()

	db.Client.TasksAPI().CreateTask(ctx, &domain.Task{})

	dataBytes, err := assets.ReadFile(niftyAsset)
	if err != nil {
		return fmt.Errorf("failed to read file %q", niftyAsset)
	}

	dataString := string(dataBytes)

	lines := strings.Split(dataString, "\n")
	for _, line := range lines {
		// Sleep for 1 minute
		time.Sleep(1 * time.Minute)
		l := strings.Split(line, ",")
		if len(l) > 6 {
			open, _ := strconv.ParseFloat(l[3], 64)
			high, _ := strconv.ParseFloat(l[4], 64)
			low, _ := strconv.ParseFloat(l[5], 64)
			close, _ := strconv.ParseFloat(l[6], 64)

			// keep string to use with techan

			_, err := parseTime(l[1], l[2])
			if err != nil {
				Logger.Error("failed to parse time", zap.Error(err))
				return err
			}

			lastPriceList := []float64{open, high, low, close}

			for _, price := range lastPriceList {
				fields := map[string]interface{}{
					"LastPrice": price,
				}

				Logger.Info("Tick", zap.Any("Last Price", fields["LastPrice"]))
				err = db.WriteData(bucket, measurement, tag, fields, time.Now())
				if err != nil {
					Logger.Error("failed to write point", zap.Any("Last Price", fields["LastPrice"]), zap.Error(err))
				}
				time.Sleep(5 * time.Second)
			}

		}
	}

	return nil
}

func parseTime(d, t string) (time.Time, error) {
	// Time format in the text file is 20210101,09:08
	// We have to convert it to 20210101090800
	parsedtime, err := utils.ToTime(fmt.Sprintf("%s%s%s", d, strings.ReplaceAll(t, ":", ""), "00"))
	return parsedtime, err
}
