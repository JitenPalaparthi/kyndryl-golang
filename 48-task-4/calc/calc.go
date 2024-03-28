package calc

import (
	"errors"
	"reflect"
)

// the returns can be float64 , since that is what the result is in or any.So purposely one function
// is given float64 return and the other any
func Add(a, b any) (any, error) {
	var sum float64

	if reflect.TypeOf(a).Name() == "string" && reflect.TypeOf(b).Name() == "string" {
		return a.(string) + b.(string), nil
	}

	switch a.(type) {
	case int:
		sum += float64(a.(int))
	case uint:
		sum += float64(a.(uint))
	case uint8:
		sum += float64(a.(uint8))
	case uint16:
		sum += float64(a.(uint16))
	case uint32:
		sum += float64(a.(uint32))
	case uint64:
		sum += float64(a.(uint64))
	case int8:
		sum += float64(a.(int8))
	case int16:
		sum += float64(a.(int16))
	case int32:
		sum += float64(a.(int32))
	case int64:
		sum += float64(a.(int64))
	case float32:
		sum += float64(a.(float32))
	case float64:
		sum += a.(float64)
	default:
		return 0, errors.New("unsupported type")
	}

	switch b.(type) {
	case int:
		sum += float64(b.(int))
	case uint:
		sum += float64(b.(uint))
	case uint8:
		sum += float64(b.(uint8))
	case uint16:
		sum += float64(b.(uint16))
	case uint32:
		sum += float64(b.(uint32))
	case uint64:
		sum += float64(b.(uint64))
	case int8:
		sum += float64(b.(int8))
	case int16:
		sum += float64(b.(int16))
	case int32:
		sum += float64(b.(int32))
	case int64:
		sum += float64(b.(int64))
	case float32:
		sum += float64(b.(float32))
	case float64:
		sum += b.(float64)
	default:
		return 0, errors.New("unsupported type")
	}
	return sum, nil
}

func Sub(a, b any) (any, error) {
	var sub float64
	switch a := a.(type) {
	case int:
		sub -= float64(a)
	case uint:
		sub -= float64(a)
	case uint8:
		sub -= float64(a)
	case uint16:
		sub -= float64(a)
	case uint32:
		sub -= float64(a)
	case uint64:
		sub -= float64(a)
	case int8:
		sub -= float64(a)
	case int16:
		sub -= float64(a)
	case int32:
		sub -= float64(a)
	case int64:
		sub -= float64(a)
	case float32:
		sub -= float64(a)
	case float64:
		sub -= a
	default:
		return 0, errors.New("unsupported type")
	}

	switch b := b.(type) {
	case int:
		sub -= float64(b)
	case uint:
		sub -= float64(b)
	case uint8:
		sub -= float64(b)
	case uint16:
		sub -= float64(b)
	case uint32:
		sub -= float64(b)
	case uint64:
		sub -= float64(b)
	case int8:
		sub -= float64(b)
	case int16:
		sub -= float64(b)
	case int32:
		sub -= float64(b)
	case int64:
		sub -= float64(b)
	case float32:
		sub -= float64(b)
	case float64:
		sub -= b
	default:
		return 0, errors.New("unsupported type")
	}
	return sub, nil
}

// int,uint, int8 int16,int32,int64,uint8 uint16,uint32,uint64,
// float32,float64
