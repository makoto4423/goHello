package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// arg := os.Args[1:]
	// if len(arg) == 0 {
	// 	fmt.Print("provide directory")
	// 	return
	// }

	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Print("error")
		return
	}
	var content []byte
	for _, file := range files {
		if !file.IsDir() && file.Size() != 0 {
			f, _ := os.ReadFile(file.Name())
			content = append(content, f...)
		}
	}
	_ = ioutil.WriteFile("out.txt", content, 0644)
}
