package utils

import (
	"strconv"
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

func GetUnit32(str string) (uint32, error) {
	u, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(u), nil
}
