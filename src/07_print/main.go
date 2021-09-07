package main

import (
	"fmt"
	_ "goHello/src/07_print/nothing"
)

func main() {
	word := "value"
	fmt.Printf("key is %v", word)
	tf := false
	fmt.Printf("tf is %t", tf)
}
