package fetcher

import (
	"context"
	"fmt"
	"os"

	"github.com/sp98/marketmoz/pkg/common"
	"github.com/sp98/marketmoz/pkg/db/influx"
	"github.com/sp98/marketmoz/pkg/fetcher/kite"
	"go.uber.org/zap"
)

var Logger *zap.Logger

// StartFetcher starts a new fetching process from the provided source
func StartFetcher(source, destination string) error {
	ctx := context.Background()
	switch source {
	case "file":
		startFileFetcher()
	case "kite":
		apiKey := os.Getenv(common.KITE_API_KEY)
		apiSecret := os.Getenv(common.KITE_API_SECRET)
		accessToken := os.Getenv(common.KITE_ACCESS_TOKEN)
		if apiKey == "" || accessToken == "" || apiSecret == "" {
			Logger.Error("failed to get Kite API key, Secret or request token from evn")
			return fmt.Errorf("failed to get Kite API key or request token from evn. API Key: %q. Request Token: %q", apiKey, accessToken)
		}

		k, err := kite.New(apiKey, apiSecret, accessToken, common.Subscriptions)
		if err != nil {
			return fmt.Errorf("failed to create new Kite connection client. Error %v", err)
		}

		k.Store = influx.NewDB(ctx, common.INFLUXDB_ORGANIZATION, common.INFLUXDB_URL, common.INFLUXDB_TOKEN)

		// Init buckets
		err = k.Store.InitBuckets()
		if err != nil {
			Logger.Error("failed to initialize influx db buckets", zap.Error(err))
			return fmt.Errorf("failed to initialize influx db buckets. Error %v", err)
		}

		// Init downsample tasks
		err = k.CreateDownsampleTasks()
		if err != nil {
			Logger.Error("failed to create downsample tasks", zap.Error(err))
			return fmt.Errorf("failed to create downsample tasks. Error %v", err)
		}

		k.StartKiteFetcher()

	default:
		return fmt.Errorf("invalid source type %q", source)
	}

	return nil
}
