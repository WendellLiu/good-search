package utils

import "reflect"

func MergeStructToMap(data interface{}, source map[string]interface{}) map[string]interface{} {
	elem := reflect.ValueOf(data).Elem()
	size := elem.NumField()

	for i := 0; i < size; i++ {
		field := elem.Type().Field(i).Name
		value := elem.Field(i).Interface()
		source[field] = value
	}

	return source
}
