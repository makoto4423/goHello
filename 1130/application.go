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

type duration int

func (d *duration) pretty() string {
	return fmt.Sprint(*d)
}
func main() {
	u := user{name: "ugly", yoo: "very"}
	u.ooo()
	u.change()
	u.ooo()
	d := duration(42)
	d.pretty()
}
