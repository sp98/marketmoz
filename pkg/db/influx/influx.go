package influx

import (
	"context"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

type DB struct {
	Organization string
	Bucket       string
	Client       influxdb2.Client
}

func NewDB(org, bucket, url, token string) *DB {
	client := influxdb2.NewClient(url, token)
	return &DB{
		Organization: org,
		Bucket:       bucket,
		Client:       client,
	}
}

func (db DB) WriteFileData(measurement string, tags map[string]string,
	fields map[string]interface{}, t time.Time) error {
	writeAPI := db.Client.WriteAPIBlocking(db.Organization, db.Bucket)
	p := influxdb2.NewPoint(measurement, tags, fields, t)
	err := writeAPI.WritePoint(context.Background(), p)
	if err != nil {
		return err
	}

	return nil
}

func (db DB) CloseDB() {
	db.Client.Close()
}
