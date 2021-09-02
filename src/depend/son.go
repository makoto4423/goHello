package depend

import (
	"fmt"
)

type Son struct {
}

func (s Son) print(ss string) {
	fmt.Println(ss)
}

func init() {
	f := new(Father)
	f.print("son call you")
	fmt.Println("now is son")
}
