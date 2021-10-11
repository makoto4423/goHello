package main

import "fmt"

func main() {

	c := make(map[string]string, 1)
	c["key"] = "val"
	fmt.Println(len(c))

	// _ = map[string]int{"key": 123}

	for k, v := range c {
		fmt.Print("key is " + k + " , value is " + v)
	}

}
