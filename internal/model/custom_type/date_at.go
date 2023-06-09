package ct

import (
	"fmt"
	"time"
)

// don't import any third party lib
const (
	daLayout = "2006-01-02"
)

type DateAt string

func (t DateAt) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%v\"", t)), nil
}

func (t *DateAt) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	tTime := (value.(time.Time)).Format(daLayout)
	*t = DateAt(tTime)
	return nil
}
