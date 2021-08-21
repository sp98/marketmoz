package influx

import (
	"context"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

type DB struct {
	Context      context.Context
	Organization string
	Client       influxdb2.Client
}

func NewDB(ctx context.Context, org, url, token string) *DB {
	client := influxdb2.NewClient(url, token)
	return &DB{
		Context:      ctx,
		Organization: org,
		Client:       client,
	}
}

func (db DB) WriteFileData(bucket, measurement string, tags map[string]string,
	fields map[string]interface{}, t time.Time) error {
	writeAPI := db.Client.WriteAPIBlocking(db.Organization, bucket)
	p := influxdb2.NewPoint(measurement, tags, fields, t)
	err := writeAPI.WritePoint(context.Background(), p)
	if err != nil {
		return err
	}

	return nil
}

func (db DB) DeleteBucket(bucket string) error {
	return db.Client.DeleteAPI().DeleteWithName(db.Context, db.Organization,
		bucket, time.Now().AddDate(-1, 0, 0), time.Now(), "")
}

func (db DB) CloseDB() {
	db.Client.Close()
}
