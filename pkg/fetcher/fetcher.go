package fetcher

import (
	"fmt"
	"os"

	"github.com/sp98/marketmoz/pkg/common"
	"github.com/sp98/marketmoz/pkg/fetcher/kite"
	"go.uber.org/zap"
)

var Logger *zap.Logger

// StartFetcher starts a new fetching process from the provided source
func StartFetcher(source, destination string) error {
	switch source {
	case "file":
		startFileFetcher()
	case "kite":
		apiKey := os.Getenv(common.KITE_API_KEY)
		requestToken := os.Getenv(common.KITE_REQUEST_TOKEN)
		if apiKey == "" || requestToken == "" {
			return fmt.Errorf("failed to get Kite API key or request token from evn. API Key: %q. Request Token: %q", apiKey, requestToken)
		}

		k, err := kite.New(apiKey, requestToken, []uint32{})
		if err != nil {
			return fmt.Errorf("failed to create new Kite connection client. Error %v", err)
		}

		k.StartKiteFetcher()

	default:
		return fmt.Errorf("invalid source type %q", source)
	}

	return nil
}
