package utils

import (
	"strconv"

	"go.uber.org/zap"
)

var Logger *zap.Logger

// GetUnit32 converts string to Unit32
func GetUnit32(str string) (uint32, error) {
	u, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(u), nil
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
