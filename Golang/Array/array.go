package main

import (
	"fmt"
)

type Array struct {
	length int
	data   []interface{}
}

func (arr *Array) get(index int) interface{} {
	return arr.data[index]
}

func (arr *Array) push(item interface{}) int {
	arr.data = append(arr.data, item)
	arr.length++

	return arr.length
}

func (arr *Array) pop() interface{} {
	lastItem := arr.data[arr.length-1]
	arr.data = arr.data[:arr.length-1]
	arr.length--

	return lastItem
}

// prints out a string in reverse
func reverse(str string) {
	if len(str) < 0 {
		fmt.Println("Invalid length of string")
	}
	for i := len(str) - 1; i >= 0; i-- {
		fmt.Print(string(str[i]))
	}

}
