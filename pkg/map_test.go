package pkg

import (
	"reflect"
	"strings"
	"testing"
)

type testMapStruct struct {
	name     string
	input    interface{}
	expected interface{}
	mapFunc  func(interface{}) interface{}
}

var mapStub = []testMapStruct{
	{
		name:     "numbers multiplied by 2",
		input:    []int{1, 2, 3, 4, 5, 6},
		expected: []int{12, 10, 8, 6, 4, 2},
		mapFunc:  func(i interface{}) interface{} { return i.(int) * 2 },
	},
	{
		name:     "map string to upper case",
		input:    []string{"1, 2, 3, 4", "he", "lo", "there"},
		expected: []string{"THERE", "LO", "HE", "1, 2, 3, 4"},
		mapFunc:  func(i interface{}) interface{} { return strings.ToUpper(i.(string)) },
	},
}

func TestMap(t *testing.T) {
	for _, testCond := range mapStub {
		var testArray Array = Array{testCond.input}
		t.Run(testCond.name, func(t *testing.T) {
			if !reflect.DeepEqual(testCond.expected, testArray.Map(testCond.mapFunc)) {
				t.Fatalf("Map Test failed %v -- %v", testCond.expected, testArray.Map(testCond.mapFunc))
			}
		})
	}
}

//benchmarking only int input
func BenchmarkMap(b *testing.B) {
	var testArray Array = Array{mapStub[0].input}
	for i := 0; i < b.N; i++ {
		if !reflect.DeepEqual(mapStub[0].expected, testArray.Map(mapStub[0].mapFunc)) {
			b.Fatalf("Map Benchmark Test failed")
		}
	}
}
