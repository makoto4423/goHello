package main

import (
	"fmt"
	"strconv"
)

func main() {

	var (
		n   int
		err error
	)

	if n, err = strconv.Atoi("123"); err != nil {
		fmt.Println("here is err")
	}

	fmt.Println("n is ", n, ". 👻 👻 👻 - you've been shadowed ;-)")
}
