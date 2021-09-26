package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	arr := make([]int, 3)
	for index, value := range arr {
		fmt.Println(index, value)
	}

	arr2 := strings.Fields("a b c")
	fmt.Println(reflect.TypeOf(arr2))

	arr3 := []int{1, 2, 3}
	fmt.Print(reflect.TypeOf(arr3))

	arr3 = append(arr3, 4)
	fmt.Println(reflect.TypeOf(arr3), arr3)
}
