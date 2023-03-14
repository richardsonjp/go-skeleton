package ct

import (
	"fmt"
	"time"
)

// don't import any third party lib
const (
	taLayout = "2006-01-02 15:04:05"
)

type TimeAt string

func (t TimeAt) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%v\"", t)), nil
}

func (t *TimeAt) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	tTime := (value.(time.Time)).Format(taLayout)
	*t = TimeAt(tTime)
	return nil
}
