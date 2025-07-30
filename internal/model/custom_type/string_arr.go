package ct

import (
	"database/sql/driver"
	"fmt"
	"strings"
)

// StringArr is a custom type for PostgreSQL text arrays
type StringArr []string

// Scan implements the sql.Scanner interface for reading from DB
func (s *StringArr) Scan(value interface{}) error {
	if value == nil {
		*s = []string{}
		return nil
	}

	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("cannot convert %T to StringArr", value)
	}

	// Remove curly braces
	str = strings.TrimPrefix(str, "{")
	str = strings.TrimSuffix(str, "}")

	if str == "" {
		*s = []string{}
	} else {
		*s = strings.Split(str, ",")
	}
	return nil
}

// Value implements the driver.Valuer interface for writing to DB
func (s StringArr) Value() (driver.Value, error) {
	if len(s) == 0 {
		return "{}", nil
	}

	// Escape any commas or special characters if needed
	escaped := make([]string, len(s))
	for i, v := range s {
		escaped[i] = `"` + strings.ReplaceAll(v, `"`, `\"`) + `"`
	}

	return "{" + strings.Join(escaped, ",") + "}", nil
}

// GormDataType defines the type GORM should use in migrations
func (StringArr) GormDataType() string {
	return "text[]"
}
