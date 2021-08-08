package influx

import (
	"github.com/influxdata/influxdb-client-go/v2"
)

type DB struct {
	Client influxdb2.Client
}

func NewDB(url string, token string) *DB {
	client := influxdb2.NewClient(url, token)

	return &DB{
		Client: client,
	}
}

func (db DB) WriteFileData(measurement string) {

}

func (db DB) CloseDB() {
	db.Client.Close()
}
