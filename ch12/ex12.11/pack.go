/*

Exercise 12.11: Write the corresponding Pack function. Given a struct value,
Pack should return a URL incorporating the parameter values from the struct.

*/

package pack

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

func Pack(ptr interface{}) (string, error) {

	vof := reflect.ValueOf(ptr)
	var v reflect.Value
	var dataErr bool
	switch vof.Kind() {
	case reflect.Interface, reflect.Pointer:
		v = vof.Elem()
		if v.Kind() != reflect.Struct {
			dataErr = true
		}
	default:
		dataErr = true
	}

	if dataErr {
		return "", fmt.Errorf("ptr must be interface or pointer to")
	}

	params := make([]string, 0, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i)
		tag := fieldInfo.Tag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}

		f := v.Field(i)
		var val string
		switch f.Kind() {
		case reflect.Invalid:
			val = "nil"

		case reflect.Bool:
			if f.Bool() {
				val = "true"
			} else {
				val = "false"
			}

		case reflect.Float32, reflect.Float64:
			val = fmt.Sprintf("%v", f.Float())

		case reflect.Int, reflect.Int8, reflect.Int16,
			reflect.Int32, reflect.Int64:
			val = fmt.Sprintf("%d", f.Int())

		case reflect.Uint, reflect.Uint8, reflect.Uint16,
			reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			val = fmt.Sprintf("%d", f.Uint())

		case reflect.String:
			val = fmt.Sprintf("\"%s\"", url.QueryEscape(f.String()))

		default:
			continue

		}
		params = append(params, fmt.Sprintf("%s=%s", name, val))
	}

	return strings.Join(params, "&"), nil
}
