package kite

import (
	"fmt"

	domain "github.com/influxdata/influxdb-client-go/v2/domain"
	"github.com/sp98/marketmoz/assets"
	"github.com/sp98/marketmoz/pkg/common"
	"github.com/sp98/marketmoz/pkg/data"
	"github.com/sp98/marketmoz/pkg/utils"
	"go.uber.org/zap"
)

var (
	ohlcDownSampleTaskAsset = "tasks/ohlc-ds.flux"
	downsamplePeriods       = []string{"1m", "3m", "5m", "10m"}
)

func GetOHLCDownSamplingTasks() (*[]domain.Task, error) {
	ohlcTasks := []domain.Task{}
	tokens := data.GetInstrumentMap()

	for tokenID, tokenDetail := range *tokens {

		for _, dsPeriod := range downsamplePeriods {
			taskName := fmt.Sprintf("OHLC-%s-%s", tokenID, dsPeriod)
			inputBucket := fmt.Sprintf(common.REAL_TIME_DATA_BUCKET, tokenDetail.Exchange, tokenDetail.Segment)
			id, err := utils.GetUnit32(tokenID)
			if err != nil {
				Logger.Error("failed to convert string to unit 32", zap.String("input", tokenID), zap.Error(err))
				continue
			}
			inputMeasurement := fmt.Sprintf(common.REAL_TIME_DATA_MEASUREMENT, id)

			outputBucket := fmt.Sprintf(common.OHLC_DOWNSAMPLE_BUCKET, tokenDetail.Exchange, tokenDetail.Segment, dsPeriod)
			outputMeasurement := fmt.Sprintf(common.OHLC_DOWNSAMPLE_MEASUREMENT, tokenID)

			dsTaskBytes, err := assets.ReadFileAndReplace(
				ohlcDownSampleTaskAsset,
				[]string{
					"${INPUT_BUCKET}", inputBucket,
					"${INPUT_MEASUREMENT}", inputMeasurement,
					"${OUTPUT_BUCKET}", outputBucket,
					"${OUTPUT_MEASUREMENT}", outputMeasurement,
				},
			)

			if err != nil {
				Logger.Error("failed to get data from task file", zap.String("filename", ohlcDownSampleTaskAsset), zap.Error(err))
				return &ohlcTasks, fmt.Errorf("failed to get data fro task file %q", ohlcDownSampleTaskAsset)
			}

			status := domain.TaskStatusTypeActive
			task := domain.Task{
				Name:   taskName,
				Every:  &dsPeriod,
				Flux:   string(dsTaskBytes),
				Status: &status,
				OrgID:  common.INFLUXDB_ORGANIZATION_ID,
			}

			ohlcTasks = append(ohlcTasks, task)
		}
	}

	return &ohlcTasks, nil
}
