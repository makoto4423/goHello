package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type collection []string
type collectionArr [2]string

func main() {
	data := collection{"a", "b"}
	change(data)
	fmt.Print(data)
	changeSlice(data)
	fmt.Print(data)
	arr := collectionArr{"a", "b"}
	changeArr(arr)
	fmt.Print(unsafe.Sizeof(arr), unsafe.Sizeof(data))
	fmt.Println(cap(arr))
}

func change(data collection) {
	data[1] = "change"
}

func changeSlice(data []string) {
	data[1] = "slice"
}

func changeArr(data [2]string) {
	data[1] = "arr"
}

func print() {
	arr := [...]string{"a", "b", "c"}
	fmt.Print(reflect.TypeOf(arr))

	arr2 := make([]string, 4)
	fmt.Print(reflect.TypeOf(arr2))

	type (
		// integer int

		bookcase [5]int
		cabinet  [5]int
		//          ^- try changing this to: integer
		//             but first: uncomment the integer type above
	)

	blue := bookcase{6, 9, 3, 2, 1}
	red := cabinet{6, 9, 3, 2, 1}

	fmt.Print("Are they equal? ")

	if cabinet(blue) == red {
		fmt.Println("✅")
	} else {
		fmt.Println("❌")
	}

	fmt.Printf("blue: %#v\n", blue)
	fmt.Printf("red : %#v\n", bookcase(red))

	arr3 := [2]int{1}
	arr4 := [2]int{1}
	fmt.Print(arr3 == arr4)

	arr5 := arr3[:1]
	fmt.Println(reflect.TypeOf(arr5))

	var arr6 [0]int
	fmt.Print(reflect.TypeOf(arr6))
	var arr7 []int
	fmt.Print(reflect.TypeOf(arr7))
}
