package kite

import (
	"fmt"
	"os"

	domain "github.com/influxdata/influxdb-client-go/v2/domain"
	"github.com/sp98/marketmoz/assets"
	"github.com/sp98/marketmoz/pkg/common"
	"github.com/sp98/marketmoz/pkg/data"
	"go.uber.org/zap"
)

var (
	ohlcDownSampleTaskAsset = "tasks/ohlc-ds.flux"
)

func GetOHLCDownSamplingTasks() (*[]domain.Task, error) {
	orgID := os.Getenv(common.INFLUXDB_ORGANIZATION_ID)
	if orgID == "" {
		return nil, fmt.Errorf("failed to get organization ID using env variable %s", common.INFLUXDB_ORGANIZATION_ID)
	}
	ohlcTasks := []domain.Task{}

	for _, sub := range common.Subscriptions {
		instrumentID := fmt.Sprintf("%d", sub)
		instrumentDetail := data.GetInstrumentDetails(fmt.Sprintf("%d", sub))
		for _, dsPeriod := range common.DownsamplePeriods {
			taskName := fmt.Sprintf("OHLC-%s-%s", instrumentID, dsPeriod)
			inputBucket, err := GetInputBucket(*instrumentDetail, dsPeriod)
			if err != nil {
				Logger.Error("failed to get input bucket to downsample instrument", zap.String("instrument", instrumentDetail.Name))
				return nil, fmt.Errorf("failed to get input bucket to downsample %s instrument", instrumentDetail.Name)
			}
			inputMeasurement, err := GetInputMeasurement(sub, dsPeriod)
			if err != nil {
				Logger.Error("failed to get input measurement to downsample instrument", zap.String("instrument", instrumentDetail.Name))
				return nil, fmt.Errorf("failed to get input measurement to downsample %s instrument", instrumentDetail.Name)
			}

			outputBucket := fmt.Sprintf(common.OHLC_DOWNSAMPLE_BUCKET, instrumentDetail.InstrumentType, instrumentDetail.Segment, instrumentDetail.Exchange, dsPeriod)
			outputMeasurement := fmt.Sprintf(common.OHLC_DOWNSAMPLE_MEASUREMENT, instrumentID)

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

			p := dsPeriod
			status := domain.TaskStatusTypeActive
			task := domain.Task{
				Name:   taskName,
				Every:  &p,
				Flux:   string(dsTaskBytes),
				Status: &status,
				OrgID:  orgID,
			}

			ohlcTasks = append(ohlcTasks, task)
		}
	}

	return &ohlcTasks, nil
}

func GetInputBucket(instrument data.Instrument, dsPeriod string) (string, error) {
	switch dsPeriod {
	case "1m":
		return fmt.Sprintf(common.REAL_TIME_DATA_BUCKET, instrument.InstrumentType, instrument.Segment, instrument.Exchange), nil
	case "5m":
		return fmt.Sprintf(common.OHLC_DOWNSAMPLE_BUCKET, instrument.InstrumentType, instrument.Segment, instrument.Exchange, "1m"), nil
	case "1d":
		return fmt.Sprintf(common.OHLC_DOWNSAMPLE_BUCKET, instrument.InstrumentType, instrument.Segment, instrument.Exchange, "5m"), nil
	}

	return "", fmt.Errorf("invalid downsample period %q", dsPeriod)
}

func GetInputMeasurement(tokenID uint32, dsPeriod string) (string, error) {
	switch dsPeriod {
	case "1m":
		return fmt.Sprintf(common.REAL_TIME_DATA_MEASUREMENT, tokenID), nil
	case "5m":
		return fmt.Sprintf(common.OHLC_DOWNSAMPLE_MEASUREMENT, fmt.Sprintf("%d", tokenID)), nil
	case "1d":
		return fmt.Sprintf(common.OHLC_DOWNSAMPLE_MEASUREMENT, fmt.Sprintf("%d", tokenID)), nil
	}

	return "", fmt.Errorf("invalid downsample period %q", dsPeriod)
}
