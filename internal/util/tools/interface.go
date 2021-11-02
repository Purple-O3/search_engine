package tools

import (
	"errors"
	"reflect"
	"time"
)

func InterfaceZero(i interface{}) bool {
	return reflect.DeepEqual(i, reflect.Zero(reflect.TypeOf(i)).Interface())
}

func InterfaceLt(a interface{}, b interface{}) (bool, error) {
	aValue := reflect.ValueOf(a)
	bValue := reflect.ValueOf(b)
	if a == nil || b == nil || aValue.Type() != bValue.Type() {
		return false, errors.New("two compare type not equal")
	}
	switch a.(type) {
	case int:
		return a.(int) < b.(int), nil
	case int8:
		return a.(int8) < b.(int8), nil
	case int16:
		return a.(int16) < b.(int16), nil
	case int32:
		return a.(int32) < b.(int32), nil
	case int64:
		return a.(int64) < b.(int64), nil

	case uint:
		return a.(uint) < b.(uint), nil
	case uint8:
		return a.(uint8) < b.(uint8), nil
	case uint16:
		return a.(uint16) < b.(uint16), nil
	case uint32:
		return a.(uint32) < b.(uint32), nil
	case uint64:
		return a.(uint64) < b.(uint64), nil

	case float32:
		return a.(float32) < b.(float32), nil
	case float64:
		return a.(float64) < b.(float64), nil

	case string:
		return a.(string) < b.(string), nil

	case time.Time:
		return a.(time.Time).Before(b.(time.Time)), nil

	default:
		return false, errors.New("not default type")
	}
}

func InterfaceGt(a interface{}, b interface{}) (bool, error) {
	aValue := reflect.ValueOf(a)
	bValue := reflect.ValueOf(b)
	if a == nil || b == nil || aValue.Type() != bValue.Type() {
		return false, errors.New("two compare type not equal")
	}
	switch a.(type) {
	case int:
		return a.(int) > b.(int), nil
	case int8:
		return a.(int8) > b.(int8), nil
	case int16:
		return a.(int16) > b.(int16), nil
	case int32:
		return a.(int32) > b.(int32), nil
	case int64:
		return a.(int64) > b.(int64), nil

	case uint:
		return a.(uint) > b.(uint), nil
	case uint8:
		return a.(uint8) > b.(uint8), nil
	case uint16:
		return a.(uint16) > b.(uint16), nil
	case uint32:
		return a.(uint32) > b.(uint32), nil
	case uint64:
		return a.(uint64) > b.(uint64), nil

	case float32:
		return a.(float32) > b.(float32), nil
	case float64:
		return a.(float64) > b.(float64), nil

	case string:
		return a.(string) > b.(string), nil

	case time.Time:
		return a.(time.Time).After(b.(time.Time)), nil

	default:
		return false, errors.New("not default type")
	}
}

func InterfaceEq(a interface{}, b interface{}) (bool, error) {
	aValue := reflect.ValueOf(a)
	bValue := reflect.ValueOf(b)
	if a == nil || b == nil || aValue.Type() != bValue.Type() {
		return false, errors.New("two compare type not equal")
	}
	switch a.(type) {
	case int:
		return a.(int) == b.(int), nil
	case int8:
		return a.(int8) == b.(int8), nil
	case int16:
		return a.(int16) == b.(int16), nil
	case int32:
		return a.(int32) == b.(int32), nil
	case int64:
		return a.(int64) == b.(int64), nil

	case uint:
		return a.(uint) == b.(uint), nil
	case uint8:
		return a.(uint8) == b.(uint8), nil
	case uint16:
		return a.(uint16) == b.(uint16), nil
	case uint32:
		return a.(uint32) == b.(uint32), nil
	case uint64:
		return a.(uint64) == b.(uint64), nil

	case float32:
		return a.(float32) == b.(float32), nil
	case float64:
		return a.(float64) == b.(float64), nil

	case string:
		return a.(string) == b.(string), nil

	case time.Time:
		return a.(time.Time).Equal(b.(time.Time)), nil

	default:
		return false, errors.New("not default type")
	}
}
