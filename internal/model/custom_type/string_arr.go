package ct

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// don't import any third party lib

type StringArr []string

// Value for sql.Valuer
func (model StringArr) Value() (driver.Value, error) {
	data, _ := json.Marshal(model)
	return data, nil
}

func (model *StringArr) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	bytes, ok := value.([]uint8)
	if !ok {
		return errors.New("failed to unmarshal StringArr value")
	}

	err := json.Unmarshal(bytes, &model)
	if err != nil {
		return err
	}
	return nil
}
