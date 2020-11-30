package main

import (
	"fmt"
)

type user struct {
	name string
	yoo  string
}

func (u user) ooo() {
	fmt.Println(u.name, u.yoo)
}

func (u *user) change() {
	u.yoo = "no"
}

func main() {
	u := user{name: "ugly", yoo: "very"}
	u.ooo()
	u.change()
	u.ooo()
}
