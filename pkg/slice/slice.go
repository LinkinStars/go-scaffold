package slice

import "reflect"

type MapKey interface {
	string | int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

// ConvertToInterfaceSlice convert slice to interface slice
func ConvertToInterfaceSlice(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("not a slice")
	}

	ret := make([]interface{}, s.Len())
	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
}

// ConvertToMap convert slice to map
func ConvertToMap[T MapKey](
	s []T) map[T]bool {
	m := make(map[T]bool)
	for _, v := range s {
		m[v] = true
	}
	return m
}

// ConvertToMapWithFunc convert slice to map with func
func ConvertToMapWithFunc[T any, K MapKey](s []T, f func(T) K) map[K]T {
	m := make(map[K]T)
	for _, v := range s {
		m[f(v)] = v
	}
	return m
}
