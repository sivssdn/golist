package main

import (
	"fmt"

	p "github.com/sivssdn/collections/pkg"
)

type Array = p.Array

func main() {
	var testIndex Array = Array{[]string{"e1", "e2", "e32", "e14"}}
	fmt.Println(testIndex.IndexOf("e2"))

	//filter example
	fmt.Println(testIndex.Filter(func(i interface{}) bool {
		return len(i.(string)) > 2
	}))

	testIndex = Array{[]int{1, 12, 300, 50}}
	fmt.Println(testIndex.IndexOf(50))

	//filter example
	fmt.Println(testIndex.Filter(func(i interface{}) bool {
		return i.(int)%2 == 0
	}))
}

//TODO:: Implement with generics
