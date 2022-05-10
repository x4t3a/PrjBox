package fields

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type (
	Fields    map[string]interface{}
	FieldsStr map[string]string

	Transformable interface {
		Transform() string
	}
)

const (
	FieldNameStatus    = "status"
	FieldNameCreatedAt = "created_at"
)

func NewFieldsAuto() Fields {
	fields := make(Fields)
	fields[FieldNameCreatedAt] = time.Now().Format(time.RFC3339)
	fields[FieldNameStatus] = FieldStatusOpen
	return fields
}

type Transofrmator func(i any) string

var transformators = map[string]Transofrmator{
	FieldNameStatus:    func(i any) string { return FieldStatusToString(i) },
	FieldNameCreatedAt: func(i any) string { return i.(string) },
}

func (f Fields) Transform() FieldsStr {
	r := make(FieldsStr, len(f))
	for k, v := range f {
		r[k] = transformators[k](v)
	}

	return r
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
