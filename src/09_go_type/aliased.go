package main

import (
	"fmt"
	"math"
)

// type 定义，= 代表赋予别名， 不加= 代表定义一个新的类型，类似于java里的实现一个接口但是不做任何事情
type alia = uint8

type alia2 uint8

func main() {
	var (
		// type byte = int8
		b byte
		i uint8
	)
	i = b
	_ = i
	c := alia(1)
	i = c
	// d := alia2(1)
	// i = d
	var a int64 = math.MaxInt64
	fmt.Println(a, int8(a))
}
