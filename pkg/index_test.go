package pkg

import "testing"

type testTable struct {
	name     string
	values   []string
	search   string
	expected int
}

//for coverage:
// go test -coverprofile=c.out
// go tool cover -func=c.out
//go tool cover -html=c.out -o coverage.htm

var stringData = []testTable{
	{
		name:     "string present in list",
		values:   []string{"e1", "e2", "e3"},
		search:   "e2",
		expected: 1,
	},
	{
		name:     "string absent in list",
		values:   []string{"test", "normal", "func", "here", "not", "a", "big", "list"},
		search:   "bigO",
		expected: -1,
	},
	{
		name:     "empty string",
		values:   []string{},
		search:   "big",
		expected: -1,
	},
}

// go test -v -run '' -bench=.
func TestIndexOf(t *testing.T) {
	for _, testCond := range stringData {
		t.Run(testCond.name, func(t *testing.T) {
			var testIndex Array = Array{testCond.values}
			if testCond.expected != testIndex.IndexOf(testCond.search) {
				t.Fatalf("Expected output for test named %v is %v", testCond.name, testCond.expected)
			}
		})
	}

}

//go test -bench=.
func BenchmarkIndexOf(b *testing.B) {
	var testIndex Array = Array{stringData[0].values}
	for i := 0; i < b.N; i++ {
		if stringData[0].expected != testIndex.IndexOf(stringData[0].search) {
			b.Fatalf("Expected output for test named %v is %v", stringData[0].name, stringData[0].expected)
		}
	}
}
