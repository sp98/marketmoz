package influx

import (
	"context"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	domain "github.com/influxdata/influxdb-client-go/v2/domain"
	"go.uber.org/zap"
)

type DB struct {
	Context      context.Context
	Organization string
	Client       influxdb2.Client
}

var Logger *zap.Logger

func NewDB(ctx context.Context, org, url, token string) *DB {
	client := influxdb2.NewClient(url, token)
	return &DB{
		Context:      ctx,
		Organization: org,
		Client:       client,
	}
}

func (db DB) WriteData(bucket, measurement string, tags map[string]string,
	fields map[string]interface{}, t time.Time) error {
	writeAPI := db.Client.WriteAPIBlocking(db.Organization, bucket)
	p := influxdb2.NewPoint(measurement, tags, fields, t)
	err := writeAPI.WritePoint(context.Background(), p)
	if err != nil {
		return err
	}
	return nil
}

func (db DB) GetData(query string) (*api.QueryTableResult, error) {
	queryAPI := db.Client.QueryAPI(db.Organization)
	result, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (db DB) WriteTask(task *domain.Task) (*domain.Task, error) {
	return db.Client.TasksAPI().CreateTask(db.Context, task)
}

func (db DB) DeleteBucket(bucket string) error {
	return db.Client.DeleteAPI().DeleteWithName(db.Context, db.Organization,
		bucket, time.Now().AddDate(-1, 0, 0), time.Now(), "")
}

func (db DB) CloseDB() {
	db.Client.Close()
}
