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

func StartTimeAndLoc() (time.Time, *time.Location) {
	now := time.Now()
	yyyy, mm, dd := now.Date()
	return time.Date(yyyy, mm, dd, 9, 35, 2, 0, now.Location()), now.Location()
}

func CurrentTime() string {
	loc, _ := time.LoadLocation("Asia/Kolkata")
	return time.Now().In(loc).Format("2006-01-02 15:04:05")
}
