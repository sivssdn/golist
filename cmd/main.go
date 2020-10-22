package main

import (
	"fmt"
	"strings"

	p "github.com/sivssdn/golist/pkg"
)

type Array = p.Array

func main() {
	var test1 Array = Array{[]string{"e1", "e2", "e32", "e14"}}
	//indexOf example
	fmt.Println(test1.IndexOf("e2"))
	//filter example
	fmt.Println(test1.Filter(func(i interface{}) bool {
		return len(i.(string)) > 2
	}))
	//equals example
	fmt.Println(test1.Equals([]string{"e1", "e2", "e14"}))
	//map example
	fmt.Println(test1.Map(func(i interface{}) interface{} {
		return strings.ToUpper(i.(string))
	}))

	test2 := Array{[]int{1, 12, 300, 50}}
	//indexOf example
	fmt.Println(test2.IndexOf(50))
	//filter example
	fmt.Println(test2.Filter(func(i interface{}) bool {
		return i.(int)%2 == 0
	}))
	//equals example
	fmt.Println(test2.Equals([]int{1, 12, 300, 50}))
	//map example
	fmt.Println(test2.Map(func(i interface{}) interface{} {
		return i.(int) * 2
	}))
}

//TODO:: Implement with generics
