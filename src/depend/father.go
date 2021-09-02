package depend

import "fmt"

func init() {
	fmt.Println("father")
	s := new(Son)
	s.print("father call son")
}

type Father struct {
}

func (f Father) print(s string) {
	fmt.Println(s)
}
