package timeutil

import (
	"time"
)

var (
	IDZone *time.Location
)

const (
	// ISO8601TimeDate Date
	ISO8601TimeDate = "2006-01-02"
	// ISO8601TimeWithoutSec
	ISO8601TimeWithoutZone = "2006-01-02 15:04:05"
	// ISO8601TimeDate in 24 Hour Format
	ISO8601TimeDateWithTimeZone = "2006-01-02T15:04:05+07:00"
	// ISO8601TimeDate in 12 Hour Format
	// Example case : date retrieved from DB value "1961-08-06T00:00:00+07:30"
	// "+07:30" means it's 12 hour format. If parse in 24 hour format error, try to use 12 hour format instead.
	ISO8601TimeDateWithTimeZone12HourFormat = "2006-01-02T15:04:05-07:00"
	// ISO8601TimeDateFromUTC
	ISO8601TimeDateFromUTC = "2006-01-02T15:04:05Z"
	//
	RFC850NoDayNoTimeNoZone = "02-Jan-2006"
)

func Init(defTimeZone string) {
	IDZone, _ = time.LoadLocation(defTimeZone)
}
