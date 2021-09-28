package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Print(slice[0:2:5])

	slice2 := []int{7, 8, 9, 0, 1, 2, 3}
	copy(slice, slice2)
	fmt.Println(slice)
	fmt.Println(slice2)
}
