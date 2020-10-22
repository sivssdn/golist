package pkg

import (
	"reflect"
	"testing"
)

type testFilterStruct struct {
	name       string
	input      interface{}
	expected   interface{}
	filterFunc func(interface{}) bool
}

var filterStub = []testFilterStruct{
	{
		name:       "numbers divisible by 2",
		input:      []int{1, 2, 3, 4, 5, 6},
		expected:   []int{6, 4, 2},
		filterFunc: func(i interface{}) bool { return i.(int)%2 == 0 },
	},
}

func TestFilter(t *testing.T) {
	for _, testCond := range filterStub {
		var testArray Array = Array{testCond.input}
		t.Run(testCond.name, func(t *testing.T) {
			if testCond.expected != testArray.Filter(testCond.filterFunc) {
				t.Fatalf("Test failed %v %v %v %v", testArray.Filter(testCond.filterFunc), testCond.expected, reflect.TypeOf(testCond.expected), reflect.TypeOf(testArray.Filter(testCond.filterFunc)))
			}
		})
	}
}
