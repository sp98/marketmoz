package utils

import (
	"time"

	"github.com/araddon/dateparse"
	"go.uber.org/zap"
)

var Logger *zap.Logger

// const (
// 	region = "Asia/Kolkata"
// 	layout = "20210101"
// )

func ToTime(in string) (time.Time, error) {
	t, err := dateparse.ParseLocal(in)
	if err != nil {
		Logger.Error("failed to parse time", zap.Error(err))
		return time.Time{}, err
	}
	return t, nil
}
