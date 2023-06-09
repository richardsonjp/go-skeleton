package ct

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// don't import any third party lib

type Json map[string]interface{}

func (model Json) Value() (driver.Value, error) {
	data, _ := json.Marshal(model)
	if string(data) == "null" {
		return nil, nil
	}
	return data, nil
}

func (model *Json) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	bytes, ok := value.([]uint8)
	if !ok {
		return errors.New("failed to unmarshal Json value")
	}

	err := json.Unmarshal(bytes, &model)
	if err != nil {
		return err
	}
	return nil
}
