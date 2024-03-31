/*

Exercise 12.12: Extend the field tag notation to express parameter validity
requirements. For example, a string might need to be a valid email address or
credit-card number, and an integer might need to be a valid US ZIP code. Modify
Unpack to check these requirements.

*/

// Package params provides a reflection-based parser for URL parameters.
package params

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

//!+Unpack

// Unpack populates the fields of the struct pointed to by ptr
// from the HTTP request parameters in req.
func Unpack(req *http.Request, ptr interface{}) error {
	if err := req.ParseForm(); err != nil {
		return err
	}

	type Value struct {
		v reflect.Value
		t reflect.StructTag
	}

	// Build map of fields keyed by effective name.
	fields := make(map[string]Value)
	v := reflect.ValueOf(ptr).Elem() // the struct variable
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i) // a reflect.StructField
		tag := fieldInfo.Tag           // a reflect.StructTag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		fields[name] = Value{v.Field(i), tag}
	}

	// Update struct field for each parameter in the request.
	for name, values := range req.Form {
		f := fields[name]
		if !f.v.IsValid() {
			continue // ignore unrecognized HTTP parameters
		}
		validator := f.t.Get("validator")
		for _, value := range values {
			if f.v.Kind() == reflect.Slice {
				elem := reflect.New(f.v.Type().Elem()).Elem()
				if err := populate(elem, value, validator); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
				f.v.Set(reflect.Append(f.v, elem))
			} else {
				if err := populate(f.v, value, validator); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
			}
		}
	}
	return nil
}

//!-Unpack

// !+populate
func populate(v reflect.Value, value string, validator string) error {
	switch v.Kind() {
	case reflect.String:
		switch validator {
		case "email":
			// fake email validation
			if !strings.Contains(value, "@") {
				return fmt.Errorf("validation failed: email")
			}
		case "ccn":
			// fake email validation
			if len(value) != 16 {
				return fmt.Errorf("validation failed: ccn invalid len")
			}
			_, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return fmt.Errorf("validation failed: ccn %v", err)
			}
		}

		v.SetString(value)

	case reflect.Int:

		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		if validator == "zip" {
			if i < 0 || i > 99999 {
				return fmt.Errorf("validation failed: zip")
			}
		}

		v.SetInt(i)

	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		v.SetBool(b)

	default:
		return fmt.Errorf("unsupported kind %s", v.Type())
	}
	return nil
}

//!-populate
