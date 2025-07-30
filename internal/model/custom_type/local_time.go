package ct

import (
	"database/sql/driver"
	"fmt"
	"time"
)

var AppLocation *time.Location

func init() {
	var err error
	AppLocation, err = time.LoadLocation("Asia/Jakarta") // Your timezone
	if err != nil {
		AppLocation = time.Local
	}
}

// LocalTime wraps time.Time and handles timezone-naive DB timestamps
type LocalTime struct {
	time.Time
}

// Scan implements sql.Scanner interface for reading from database
func (lt *LocalTime) Scan(value interface{}) error {
	switch v := value.(type) {
	case time.Time:
		// Database returns time without timezone, interpret as local
		lt.Time = time.Date(v.Year(), v.Month(), v.Day(),
			v.Hour(), v.Minute(), v.Second(),
			v.Nanosecond(), AppLocation)
	case string:
		// Handle string timestamps
		layouts := []string{
			"2006-01-02 15:04:05",
			"2006-01-02T15:04:05",
			"2006-01-02 15:04:05.000000",
			time.RFC3339,
		}

		var parsed time.Time
		var err error
		for _, layout := range layouts {
			parsed, err = time.ParseInLocation(layout, v, AppLocation)
			if err == nil {
				break
			}
		}
		if err != nil {
			return fmt.Errorf("cannot parse time string %q: %v", v, err)
		}
		lt.Time = parsed
	case []byte:
		return lt.Scan(string(v))
	case nil:
		lt.Time = time.Time{}
	default:
		return fmt.Errorf("cannot scan %T into LocalTime", value)
	}
	return nil
}

// Value implements driver.Valuer interface for writing to database
func (lt LocalTime) Value() (driver.Value, error) {
	if lt.Time.IsZero() {
		return nil, nil
	}
	// Store in local timezone without timezone info
	return lt.Time.In(AppLocation), nil
}

// Helper methods
func (lt LocalTime) IsZero() bool {
	return lt.Time.IsZero()
}

func (lt LocalTime) String() string {
	return lt.Time.Format("2006-01-02 15:04:05")
}

// Factory function to create LocalTime from current time
func NewLocalTime() LocalTime {
	return LocalTime{Time: time.Now().In(AppLocation)}
}

// Factory function to create LocalTime from time.Time
func NewLocalTimeFromTime(t time.Time) LocalTime {
	return LocalTime{Time: t.In(AppLocation)}
}
