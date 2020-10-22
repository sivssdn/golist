package pkg

import (
	"reflect"
)

func (list *Array) Map(transform func(x interface{}) interface{}) interface{} {
	dataType := reflect.TypeOf(list.Value)
	data := reflect.ValueOf(list.Value)
	result := reflect.MakeSlice(dataType, 0, data.Len())
	var transformedValue interface{}
	for i := data.Len() - 1; i >= 0; i-- {
		transformedValue = transform(data.Index(i).Interface())
		result = reflect.Append(result, reflect.ValueOf(transformedValue))
	}
	return result.Interface()
}
