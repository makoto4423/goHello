package main

import "fmt"

func main() {
	var arr1 [3]*string
	arr2 := [3]*string{new(string), new(string), new(string)}
	*arr2[0] = "ee"
	arr1 = arr2
	fmt.Print(arr1)
	slice := []string{"a", "b", "c", "d", "e", "f"}
	var newSlice []string
	newSlice = slice[1:3]
	newSlice = append(newSlice, "g")
	fmt.Print(newSlice[2])
}
