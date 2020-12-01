package main

import (
	"fmt"
)

type user struct {
	name string
	age  int
}

type admin struct {
	// 如果进行写成 abc user，则admin不持有user的方法，需显示标注abc的成员间接调用
	user
	level string
}

func (u user) change(s string) {
	u.name = s
}

func (u *user) changeAge(i int) {
	u.age = i
}

func main() {
	u := user{
		name: "user",
		age:  0,
	}
	u.change("cha")
	fmt.Println(u.name)
	u.changeAge(10)
	fmt.Println(u.age)
	ad := admin{
		user:  user{},
		level: "",
	}
	ad.change("")
}
