package fetcher

import (
	"fmt"

	"go.uber.org/zap"
)

var Logger *zap.Logger

// StartFetcher starts a new fetching process from the provided source
func StartFetcher(source, destination string) error {
	switch source {
	case "file":
		startFileFetcher()
	default:
		return fmt.Errorf("invalid source type %q", source)
	}

	return nil
}
