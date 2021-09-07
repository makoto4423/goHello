package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	_ "strings"
	_ "unicode/utf8"
)

func main() {
	str := os.Args[1]
	float, _ := strconv.ParseFloat(str, 64)
	fmt.Println(float)
	// fmt.Println(os.Args[0])
	// fmt.Println(strconv.FormatBool(true))

	// fmt.Println(len("中a"))
	// fmt.Println(utf8.RuneCountInString("中a"))
	// str = str + strings.Repeat("!", len(str))
	// fmt.Println(str)

	// var b byte = 255
	// fmt.Printf("%08b = %d\n", b, b)
	fmt.Println("uint64 :", 0, uint64(math.MaxUint64))

	f := float32(math.MaxFloat32)
	fmt.Println("max float    :", f*1.2)
}
