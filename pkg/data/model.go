package data

import (
	"time"

	"go.uber.org/zap/zapcore"
)

type FileData struct {
	Stock string
	Open  float64
	High  float64
	Low   float64
	Close float64
	Time  time.Time
}

func (d *FileData) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("Stock", d.Stock)
	enc.AddFloat64("Open", d.Open)
	enc.AddFloat64("High", d.High)
	enc.AddFloat64("Low", d.Low)
	enc.AddFloat64("Close", d.Close)
	enc.AddTime("Time", d.Time)
	return nil
}
