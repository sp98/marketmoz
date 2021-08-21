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
	db := influx.NewDB(ctx, common.ORGANIZATION, bucket, common.INFLUXDB_URL, common.INFLUXDB_TOKEN)
	defer db.Client.Close()

	db.Client.TasksAPI().CreateTask(ctx, &domain.Task{})

	dataBytes, err := assets.ReadFile(niftyAsset)
	if err != nil {
		return fmt.Errorf("failed to read file %q", niftyAsset)
	}

	dataString := string(dataBytes)
	//dataList := []data.FileData{}
	lines := strings.Split(dataString, "\n")
	for _, line := range lines {
		// Sleep for 1 second
		time.Sleep(1 * time.Minute)
		l := strings.Split(line, ",")
		if len(l) > 6 {
			open, _ := strconv.ParseFloat(l[3], 64)
			high, _ := strconv.ParseFloat(l[4], 64)
			low, _ := strconv.ParseFloat(l[5], 64)
			close, _ := strconv.ParseFloat(l[6], 64)
			_, err := parseTime(l[1], l[2])
			if err != nil {
				Logger.Error("failed to parse time", zap.Error(err))
				return err
			}

			// d := data.FileData{
			// 	Stock: l[0],
			// 	Open:  open,
			// 	High:  high,
			// 	Low:   low,
			// 	Close: close,
			// 	Time:  t,
			// }

			lastPriceList := []float64{open, high, low, close}

			for _, price := range lastPriceList {
				fields := map[string]interface{}{
					"LastPrice": price,
				}

				Logger.Info("Tick", zap.Any("Last Price", fields["LastPrice"]))
				err = db.WriteFileData(measurement, tag, fields, time.Now())
				if err != nil {
					Logger.Error("failed to write point", zap.Any("Last Price", fields["LastPrice"]), zap.Error(err))
				}
				time.Sleep(5 * time.Second)
			}

			//dataList = append(dataList, d)
		}
	}

	//Logger.Info("Data List created", zap.Int("count", len(dataList)))

	return nil
}

func parseTime(d, t string) (time.Time, error) {
	// Time format in the text file is 20210101,09:08
	// We have to convert it to 20210101090800
	parsedtime, err := utils.ToTime(fmt.Sprintf("%s%s%s", d, strings.ReplaceAll(t, ":", ""), "00"))
	return parsedtime, err
}
