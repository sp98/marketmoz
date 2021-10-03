package common

const (
	INFLUXDB_ORGANIZATION    = "marketmoz"
	INFLUXDB_ORGANIZATION_ID = "INFLUXDB_ORGANIZATION_ID"
	INFLUXDB_URL             = "http://localhost:8086/"
	INFLUXDB_TOKEN           = "m5txwvJXRbatNQM0AYKl9gkvtWVTkt_vIKU7IWotXQ-RAA-Q3i0wRrQfJTLvDmmn0e0GkCFJ0lZ3w8Pb-O_4uA=="

	KITE_API_KEY      = "KITE_API_KEY"
	KITE_ACCESS_TOKEN = "KITE_ACCESS_TOKEN"
	KITE_API_SECRET   = "KITE_API_SECRET"

	SLACK_WEBHOOK_URL = "SLACK_WEBHOOK_URL"
	SLACK_CHANNEL     = "SLACK_MARKETMOZ_CHANNEL"

	// Query files
	OHLC_QUERY_ASSET      = "queries/ohlc.flux"
	OHLC_QUERY_TEST_ASSET = "queries/test.flux"
	LASTPRICE_QUERY_ASSET = "queries/lastPrice.flux"
)

var (
	DownsamplePeriods = []string{"1m", "5m"}
	Subscriptions     = []uint32{408065, 738561, 5633, 2513665, 492033}

	// Buckets

	// REAL_TIME_DATA_BUCKET is the bucket that stores real time data. Naming convention is
	// RTD-<type>-<segment>-<exchange>. For example:  RTD-EQ-NSE-NSE
	// Retention Time for this bucket is 20 minutes
	REAL_TIME_DATA_BUCKET = "RTD-%s-%s-%s"

	// OHLC_DOWNSAMPLE_BUCKET is the bucket containing the downsampled OHLC data for different time period.
	// Naming convention is OHLC-<type>-<segment>-<exchange>-<time>. For example: OHLC-EQ-NSE-NSE-1m
	// Retention Time:
	// 1m : 5 days
	// 5m : 25 days
	// 1d : 60 days
	OHLC_DOWNSAMPLE_BUCKET = "OHLC-%s-%s-%s-%s"

	// Measurements

	// REAL_TIME_DATA_MEASUREMENT is the measurement that stores the real time data. Naming convention is
	//RTD-<instrument-token>. For exampe:  RTD-12123
	REAL_TIME_DATA_MEASUREMENT = "RTD-%d"

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
	STOPLOSS: %s
	TIME: %s
	______________________________________

	`
)
