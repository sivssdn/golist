package pkg

import (
	"reflect"
)

//Array contains the value of input array
type Array struct {
	Value interface{}
}

const WORKERS = 2

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

func (list *Array) IndexOfParallel(input interface{}) int {
	arr := reflect.ValueOf(list.Value)
	equalPartsLength := arr.Len() / WORKERS
	guard := make(chan bool, WORKERS)
	result := make(chan int, WORKERS)

	for i := 0; i < arr.Len(); i = i + equalPartsLength {
		guard <- true
		go func(offset int, arr interface{}, inputVal interface{}) {
			result <- (&Array{arr}).IndexOf(inputVal) + offset
			<-guard
		}(i, arr.Slice(i, i+equalPartsLength).Interface(), input)
	}

	for i := 0; i < arr.Len(); i = i + equalPartsLength {
		if index := <-result; index != -1 {
			return index
		}
	}
	return -1
}
