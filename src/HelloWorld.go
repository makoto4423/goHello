package main

import (
	"fmt"
	_ "goHello/src/depend"
	"goHello/src/makoto"
	"log"
)

func init() {
	fmt.Println("me...")
	log.Println("init....")
}

func main() {
	//var arr1 [3]*string
	//arr2 := [3]*string{new(string), new(string), new(string)}
	//*arr2[0] = "ee"
	//arr1 = arr2
	//fmt.Print(arr1)
	//slice := []string{"a", "b", "c", "d", "e", "f"}
	//var newSlice []string
	//newSlice = slice[1:3]
	//newSlice = append(newSlice, "g")
	//fmt.Print(newSlice[2])
	// 绝了，看不出log和fmt的输出顺序关系，会随着代码的顺序不同而不同
	// fmt在最前面时，fmt一定先行，log在最前面时顺序随机？i
	fmt.Println("10")
	log.Println("start...")
	makoto.Void()
	fmt.Println(makoto.Ff(12))

	for i := 0; i < 100; i++ {
		go goFunc(i)
	}

}

func goFunc(i int) {
	fmt.Println("goroutine ", i)
}
