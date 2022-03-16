package fields

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type Fields map[string]string

func NewFieldsAuto() Fields {
	fields := make(Fields)
	fields["created_at"] = time.Now().Format(time.RFC3339)
	return fields
}

func (f Fields) Value() (driver.Value, error) {
	return json.Marshal(f)
}

func (a *Fields) Scan(value interface{}) error {
	if *a == nil {
		*a = make(Fields)
	}

	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}
