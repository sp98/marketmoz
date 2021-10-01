package influx

import (
	"context"
	"fmt"
	"strings"
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

func (db DB) InitBuckets() error {
	buckets, err := GetBuckets()
	if err != nil {
		return fmt.Errorf("failed to get buckets. Error %v", err)
	}

	for _, b := range *buckets {
		res, err := db.Client.BucketsAPI().CreateBucket(db.Context, &b)
		if err != nil {
			if strings.Contains(err.Error(), "conflict") {
				Logger.Info("Bucket already created", zap.String("bucket", b.Name))
				continue
			}
			return fmt.Errorf("failed to create bucket. Error %v", err)
		}

		Logger.Info("Successfully created bucket", zap.String("bucket", res.Name))
	}

	return nil
}

func (db DB) WriteTickData(bucket, measurement string, tags map[string]string,
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

func (db DB) FindTask(filter *api.TaskFilter) ([]domain.Task, error) {
	return db.Client.TasksAPI().FindTasks(db.Context, filter)
}

func (db DB) DeleteTask(task *domain.Task) error {
	return db.Client.TasksAPI().DeleteTask(db.Context, task)
}

func (db DB) DeleteBucket(bucket string) error {
	return db.Client.DeleteAPI().DeleteWithName(db.Context, db.Organization,
		bucket, time.Now().AddDate(-1, 0, 0), time.Now(), "")
}

func (db DB) CloseDB() {
	db.Client.Close()
}
