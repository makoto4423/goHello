package main

import (
	"fmt"
	"time"
)

func main() {
	h, _ := time.ParseDuration("4h30m")
	fmt.Println(h.Hours())

	m := 2
	h *= time.Duration(m)
	fmt.Println(h)
}
