package data

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sp98/marketmoz/pkg/common"
	"github.com/sp98/marketmoz/pkg/db/influx"
)

func Delete(organization, bucket string) error {
	ctx := context.Background()
	db := influx.NewDB(ctx, organization, common.INFLUXDB_URL, common.INFLUXDB_TOKEN)
	err := db.DeleteBucket(bucket)
	if err != nil {
		return errors.Wrapf(err, "failed to delete data in organization %q and bucket %q", organization, bucket)
	}
	return nil
}

func DeleteMeasurement(organization, bucket, measurement string) error {

	// TODO: How to delete measurements inside bucket
	// ctx := context.Background()
	// db := influx.NewDB(ctx, organization, common.INFLUXDB_URL, common.INFLUXDB_TOKEN)
	// err := db.DeleteAllData(bucket)
	// if err != nil {
	// 	return errors.Wrapf(err, "failed to delete data in organization %q and bucket %q", organization, bucket)
	// }
	return nil
}
