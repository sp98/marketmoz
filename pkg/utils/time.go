package utils

import (
	"time"

	"github.com/araddon/dateparse"
	"go.uber.org/zap"
)

// ToTime converts string to time
func ToTime(in string) (time.Time, error) {
	t, err := dateparse.ParseLocal(in)
	if err != nil {
		Logger.Error("failed to parse time", zap.Error(err))
		return time.Time{}, err
	}
	return t, nil
}
