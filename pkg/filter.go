package pkg

import (
	"reflect"
)

func (list *Array) Filter(condition func(x interface{}) bool) interface{} {
	dataType := reflect.TypeOf(list.Value)
	data := reflect.ValueOf(list.Value)
	result := reflect.MakeSlice(dataType, 0, 0)
	for i := data.Len() - 1; i >= 0; i-- {
		if condition(data.Index(i).Interface()) {
			result = reflect.Append(result, data.Index(i))
		}
	}
	return result
}
