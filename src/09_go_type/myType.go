package main

import (
	"fmt"
	"goHello/src/09_go_type/ano"
)

type (
	Neko int
	Enu  int
)

func main() {
	// 会报错，不同类型不能直接赋值转换，即使实际上是同种类型
	// neko := 10
	// neko = Neko(ano.Animal(1))
	neko := Neko(1)
	enu := 1
	neko = Neko(ano.Animal(1))
	fmt.Println(neko, enu)
	lion := Neko(ano.Animal(1))
	fmt.Println(lion)
	var (
		cat Neko
		dog Enu
	)
	cat = dog
	_ = cat
}
