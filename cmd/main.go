package main

import (
	"fmt"

	p "github.com/sivssdn/collections/pkg"
)

type Array = p.Array

func main() {
	var testIndex Array = Array{[]string{"e1", "e2", "e3"}}
	fmt.Println(testIndex.IndexOf("e2"))

	testIndex = Array{[]int{1, 12, 300, 50}}
	fmt.Println(testIndex.IndexOf(50))
}

//TODO:: Implement with generics
