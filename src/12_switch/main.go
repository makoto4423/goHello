package main

import (
	"fmt"
	"time"
)

func main() {

	i := 142

	switch {
	case i > 100:
		fmt.Print("big ")
		fallthrough
	case i > 0:
		fmt.Print("positive ")
		fallthrough
	default:
		fmt.Print("number")
	}

	switch h := time.Now().Hour(); {
	case h > 13:
		fmt.Println("艾蜜莉")
	}

	switch h := time.Now().Hour(); h {
	case 14:
		fmt.Print("回春丹")
	}
}
