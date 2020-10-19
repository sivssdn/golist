package pkg

import (
	"reflect"
)

//Array contains the value of input array
type Array struct {
	Value interface{}
}

//IndexOf returns integer index of input element in the array
func (list *Array) IndexOf(input interface{}) int {
	index := -1
	arr := reflect.ValueOf(list.Value)
	inputVal := reflect.ValueOf(input)
	for i := arr.Len() - 1; i >= 0; i-- {
		if arr.Index(i).Interface() == inputVal.Interface() {
			index = i
		}
	}
	return index
}
