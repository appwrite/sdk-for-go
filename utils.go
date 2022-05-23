package appwrite

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// ToString changes arg to string
func ToString(arg interface{}) string {
	var tmp = reflect.Indirect(reflect.ValueOf(arg)).Interface()
	switch v := tmp.(type) {
	case int:
		return strconv.Itoa(v)
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case string:
		return v
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case fmt.Stringer:
		return v.String()
	case []string:
		joined_string := strings.Join(v, `","`)
		return "[" + joined_string + "]"
	case reflect.Value:
		return ToString(v.Interface())
	default:
		return ""
	}
}

// returns if an element is contained in an array s
func contains[T comparable](s []T, element T) bool {
	var result bool = false
	for _, el := range s {
		if el == element {
			result = true
			break
		}
	}
	return result
}

type Optional[T any] struct {
	Value     T
	Specified bool
}
