package ct

import (
	"database/sql/driver"
	"errors"
	"github.com/shopspring/decimal"
	"math"
)

// Define Decimal as a wrapper around shopspring/decimal.Decimal
type Decimal decimal.Decimal

func (model Decimal) Float64() float64 {
	value, _ := decimal.Decimal(model).Float64()
	return math.Ceil(value*100) / 100
}

// Value converts Decimal to a format suitable for database storage
func (model Decimal) Value() (driver.Value, error) {
	// Convert the decimal.Decimal to a string representation
	val, _ := decimal.Decimal(model).Float64()
	return math.Ceil(val*100) / 100, nil
}

// Scan converts a database value into a Decimal
func (model *Decimal) Scan(value interface{}) error {
	if value == nil {
		// Handle the case where the value is null in the database
		*model = Decimal(decimal.Decimal{})
		return nil
	}

	var decimalValue decimal.Decimal
	var err error
	// Convert the database value to a string

	switch v := value.(type) {
	case string:
		decimalValue, err = decimal.NewFromString(v)
		if err != nil {
			return err
		}
	case float64:
		decimalValue = decimal.NewFromFloat(v)
	case int64:
		decimalValue = decimal.NewFromInt(v)
	default:
		return errors.New("invalid type for Decimal")
	}

	*model = Decimal(decimalValue)
	return nil
}
