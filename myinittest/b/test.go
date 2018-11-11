package b

import (
	"awesomeProject/myinittest/c"
	"awesomeProject/mylang"
	"fmt"
)

func init() {
	mylang.Print("b")

}

//func Dis(){
//	a.Dis()
//}

type B struct {
	Name string
}

func (b B) GetName() {
	fmt.Println(b.Name)
}

func (b B) Dis(obj interface{}) {
	fmt.Println(obj.(c.C).GetName)
}
