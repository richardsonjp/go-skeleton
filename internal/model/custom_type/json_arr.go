package ct

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// don't import any third party lib

type JsonArr []map[string]interface{}

func (model JsonArr) Value() (driver.Value, error) {
	data, _ := json.Marshal(model)

	if string(data) == "null" {
		return nil, nil
	}
	return data, nil
}

func (model *JsonArr) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	bytes, ok := value.([]uint8)
	if !ok {
		return errors.New("failed to unmarshal JsonArr value")
	}

	err := json.Unmarshal(bytes, &model)
	if err != nil {
		return err
	}
	return nil
}
