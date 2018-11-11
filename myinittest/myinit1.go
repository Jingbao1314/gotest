package myinittest

import "fmt"

type FuncTest func()

func (f FuncTest) Dis() {
	f()
}

func Play() {
	fmt.Println("xxxxxxxxx")
}

func MyFuncTest() {
	var f FuncTest = Play
	f.Dis()
}
