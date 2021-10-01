package influx

import (
	"fmt"
	"os"

	domain "github.com/influxdata/influxdb-client-go/v2/domain"
	"github.com/sp98/marketmoz/pkg/common"
)

func GetBuckets() (*[]domain.Bucket, error) {
	orgID := os.Getenv(common.INFLUXDB_ORGANIZATION_ID)
	if orgID == "" {
		return nil, fmt.Errorf("failed to get organization ID using env variable %s", common.INFLUXDB_ORGANIZATION_ID)
	}

	return &[]domain.Bucket{
		{
			Name:  "RTD-EQ-NSE-NSE",
			OrgID: &orgID,
			RetentionRules: domain.RetentionRules{
				{
					EverySeconds: 3600, // 1 hr
				},
			},
		},
		{
			Name:  "OHLC-EQ-NSE-NSE-1m",
			OrgID: &orgID,
			RetentionRules: domain.RetentionRules{
				{
					EverySeconds: 604800, // 7 days
				},
			},
		},
		{
			Name:  "OHLC-EQ-NSE-NSE-5m",
			OrgID: &orgID,
			RetentionRules: domain.RetentionRules{
				{
					EverySeconds: 1200000, // 30 days
				},
			},
		},
	}, nil
}
