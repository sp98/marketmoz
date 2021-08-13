package data

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sp98/marketmoz/pkg/common"
	"github.com/sp98/marketmoz/pkg/db/influx"
)

func Delete(organization, bucket string) error {
	ctx := context.Background()
	db := influx.NewDB(ctx, organization, bucket, common.INFLUXDB_URL, common.INFLUXDB_TOKEN)
	err := db.DeleteAllData()
	if err != nil {
		return errors.Wrapf(err, "failed to delete data in organization %q and bucket %q", organization, bucket)
	}
	return nil
}
