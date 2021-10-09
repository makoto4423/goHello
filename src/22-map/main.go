package main

import "fmt"

func main() {
	var m map[string]int
	m["key"] = 123
	c := make(map[string]string, 1)
	c["key"] = "value"
	fmt.Println(len(c))

}
