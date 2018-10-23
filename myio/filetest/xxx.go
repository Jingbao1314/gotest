package filetest

import (
	"fmt"
)

type Person struct {
	Age int
}

func (this Person) Say() {
	fmt.Println(7721)
}

func FuncTest() {
	person := Person{11}
	person.Say()
}
