package a

import (
	"awesomeProject/myinittest/c"
	"fmt"
)

func init() {
	fmt.Println("a")
}

//func Dis()  {
//	fmt.Println("dis a")
//
//}

type A struct {
	Name string
}

func (a A) GetName() {
	fmt.Println(a.Name)
}

func (a A) Dis(obj interface{}) {
	obj.(c.C).GetName()
}
