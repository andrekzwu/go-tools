package util

import (
	"reflect"
	"regexp"
	"strings"
)

// TrimSpace
func TrimSpace(data interface{}) reflect.Value {
	return trimSpace(reflect.ValueOf(data))
}

// trimSpace
func trimSpace(value reflect.Value) reflect.Value {
	switch t := value.Kind(); t {
	case reflect.String:
		value = trimString(value)
	case reflect.Ptr:
		value = trimSpace(value.Elem())
	case reflect.Struct:
		value = trimStruct(value)
	case reflect.Slice:
		value = trimSlice(value)
	case reflect.Interface:
		value = trimSpace(reflect.ValueOf(value.Interface()))
	default:

	}
	return value
}

// trimString
func trimString(value reflect.Value) reflect.Value {
	if value.CanSet() {
		value.Set(reflect.ValueOf(strings.TrimSpace(value.Interface().(string))))
	} else {
		value = reflect.ValueOf(strings.TrimSpace(value.Interface().(string)))
	}
	return value
}

// trimStruct
func trimStruct(v reflect.Value) reflect.Value {
	for i := 0; i < v.NumField(); i++ {
		fieldVal := v.Field(i)
		if !fieldVal.CanSet() {
			continue
		}
		fieldVal = trimSpace(fieldVal)
		if v.Field(i).Kind() == reflect.Ptr {
			fieldVal = fieldVal.Addr()
		}
		v.Field(i).Set(fieldVal)
	}
	return v
}

// trimSlice
func trimSlice(v reflect.Value) reflect.Value {
	for i := 0; i < v.Len(); i++ {
		val := v.Index(i)
		if !val.CanSet() {
			continue
		}
		val = trimSpace(val)
		if v.Index(i).Kind() == reflect.Ptr {
			val = val.Addr()
		}
		v.Index(i).Set(val)
	}
	return v
}

// compress string
func CompressStr(str, repl string) string {
	if str == "" {
		return ""
	}
	reg := regexp.MustCompile("\\s+")
	return reg.ReplaceAllString(str, repl)
}
