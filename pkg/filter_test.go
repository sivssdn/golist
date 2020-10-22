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
	{
		name:       "string length greater than 2",
		input:      []string{"1, 2, 3, 4", "he", "lo", "there"},
		expected:   []string{"there", "1, 2, 3, 4"},
		filterFunc: func(i interface{}) bool { return len(i.(string)) > 2 },
	},
}

func TestFilter(t *testing.T) {
	for _, testCond := range filterStub {
		var testArray Array = Array{testCond.input}
		t.Run(testCond.name, func(t *testing.T) {
			if !reflect.DeepEqual(testCond.expected, testArray.Filter(testCond.filterFunc)) {
				t.Fatal("Filter Test failed")
			}
		})
	}
}

//benchmarking only int input
func BenchmarkFilter(b *testing.B) {
	var testArray Array = Array{filterStub[0].input}
	for i := 0; i < b.N; i++ {
		if !reflect.DeepEqual(filterStub[0].expected, testArray.Filter(filterStub[0].filterFunc)) {
			b.Fatal("Filter Benchmark Test failed")
		}
	}
}
