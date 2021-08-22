package influx

import (
	"fmt"

	domain "github.com/influxdata/influxdb-client-go/v2/domain"
	"github.com/sp98/marketmoz/assets"
	"github.com/sp98/marketmoz/pkg/common"
	"go.uber.org/zap"
)

var (
	ohlcDownSampleTaskAsset = "/scripts/ohlc-ds.flux"
	downsamplePeriods       = []string{"1m", "3m", "5m", "10m"}
)

func GetOHLCDownSamplingTasks() *[]domain.Task {
	ohlcTasks := []domain.Task{}
	tokens := common.GetTokenMap()

	for tokenID, tokenDetail := range *tokens {

		for _, dsPeriod := range downsamplePeriods {
			taskName := fmt.Sprintf("OHLC-%s-%s", tokenID, dsPeriod)
			fromBucket := fmt.Sprintf(common.REAL_TIME_DATA_BUCKET, tokenDetail.Exchange, tokenDetail.Segment)
			fromMeasurement := fmt.Sprintf(common.REAL_TIME_DATA_MEASUREMENT, tokenID)

			toBucket := fmt.Sprintf(common.OHLC_DOWNSAMPLE_BUCKET, tokenDetail.Exchange, tokenDetail.Segment, dsPeriod)
			toMeasurement := fmt.Sprintf(common.OHLC_DOWNSAMPLE_MEASUREMENT, tokenID)

			dsTaskBytes, err := assets.ReadFileAndReplace(
				ohlcDownSampleTaskAsset,
				[]string{
					"${TASK_NAME}", taskName,
					"${FROM_BUCKET}", fromBucket,
					"${FROM_MEASUREMENT}", fromMeasurement,
					"${TO_BUCKET}", toBucket,
					"${TO_MEASUREMENT}", toMeasurement,
					"${DS_TIME}", dsPeriod,
				},
			)

			if err != nil {
				Logger.Error("failed to get data from task file", zap.String("filename", ohlcDownSampleTaskAsset))
				continue
			}

			status := domain.TaskStatusTypeActive
			task := domain.Task{
				Name:   taskName,
				Every:  &dsPeriod,
				Flux:   string(dsTaskBytes),
				Status: &status,
			}

			ohlcTasks = append(ohlcTasks, task)
		}
	}

	return &ohlcTasks
}
