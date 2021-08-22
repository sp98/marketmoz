package common

const (
	INFLUXDB_ORGANIZATION = "marketmoz"
	INFLUXDB_URL          = "http://localhost:8086/"
	INFLUXDB_TOKEN        = "m5txwvJXRbatNQM0AYKl9gkvtWVTkt_vIKU7IWotXQ-RAA-Q3i0wRrQfJTLvDmmn0e0GkCFJ0lZ3w8Pb-O_4uA=="

	KITE_API_KEY       = "KITE_API_KEY"
	KITE_REQUEST_TOKEN = "KITE_REQUEST_TOKEN"
)

var (
	// Buckets

	// REAL_TIME_DATA_BUCKET is the bucket that stores real time data. Naming convention is
	// <exchange>-<segment>-RTD. For example:  NSE-EQ-RTD, BSE-EQ-RTD
	REAL_TIME_DATA_BUCKET = "%s-%s-RTD"

	// Measurements

	// REAL_TIME_DATA_MEASUREMENT is the measurement that stores the real time data. Naming convention is
	// <instrument-token>-RTD. For exampe:  12123-RTD
	REAL_TIME_DATA_MEASUREMENT = "%d-RTD"
)
