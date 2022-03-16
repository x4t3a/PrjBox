package multipart

import (
	"fmt"
	"net/http"
	"reflect"
)

// TODO
func Decode(v interface{}, tagName string, r *http.Request) error {
	rv := reflect.ValueOf(v)
	t := reflect.TypeOf(v)
	if t == nil {
		return fmt.Errorf("v nil")
	}
	for i := 0; i < rv.NumField(); i++ {
		field := rv.Field(i)
		if field.Kind() == reflect.Struct{
			Decode(field, tagName, r)
		} else {
			tag := t.Field(i).Tag.Get(tagName)
			if tag != "" {
				field.SetString(r.FormValue(tag))
			}
		}
	}
	
	return nil
}
