package common

const (
	INFLUXDB_ORGANIZATION    = "marketmoz"
	INFLUXDB_ORGANIZATION_ID = "b70d469d73c00531"
	INFLUXDB_URL             = "http://localhost:8086/"
	INFLUXDB_TOKEN           = "m5txwvJXRbatNQM0AYKl9gkvtWVTkt_vIKU7IWotXQ-RAA-Q3i0wRrQfJTLvDmmn0e0GkCFJ0lZ3w8Pb-O_4uA=="

	KITE_API_KEY       = "KITE_API_KEY"
	KITE_REQUEST_TOKEN = "KITE_REQUEST_TOKEN"
	KITE_API_SECRET    = "KITE_API_SECRET"

	SLACK_WEBHOOK_URL = "SLACK_WEBHOOK_URL"
	SLACK_CHANNEL     = "SLACK_MARKETMOZ_CHANNEL"

	// Query files
	OHLC_QUERY_ASSET      = "queries/ohlc.flux"
	OHLC_QUERY_TEST_ASSET = "queries/test.flux"
	LASTPRICE_QUERY_ASSET = "queries/lastPrice.flux"
)

var (
	// Buckets

	// REAL_TIME_DATA_BUCKET is the bucket that stores real time data. Naming convention is
	// <exchange>-<segment>-RTD. For example:  NSE-EQ-RTD, BSE-EQ-RTD
	REAL_TIME_DATA_BUCKET = "RTD-%s-%s"

	// Measurements

	// REAL_TIME_DATA_MEASUREMENT is the measurement that stores the real time data. Naming convention is
	// <instrument-token>-RTD. For exampe:  12123-RTD
	REAL_TIME_DATA_MEASUREMENT = "RTD-%d"

	// OHLC_DOWNSAMPLE_BUCKET is the bucket containing the downsampled OHLC data for different time period.
	// Naming convention is OHLC-<Exchange>-<SEGMENT>-<Time>. For example: OHLC-NSE-EQ-1m
	OHLC_DOWNSAMPLE_BUCKET = "OHLC-%s-%s-%s"

	// OHLC_DOWNSAMPLE_MEASUREMENT is the measurement to store downsampled OHLC data for a token.
	// Naming convention is OHLC-<Instruement-Token>. For example: OHLC-12345
	OHLC_DOWNSAMPLE_MEASUREMENT = "OHLC-%s"

	// Slack message format
	SLACK_NOTIFICATION_MESSAGE = `TRADE TYPE: %s
	STATUS: %s
	NAME: %s
	SYMBOL: %s
	EXCHANGE: %s
	SEGMENT: %s
	TIME: %s
	______________________________________

	`
)
