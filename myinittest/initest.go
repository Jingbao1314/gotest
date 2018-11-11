package myinittest

import (
	"awesomeProject/myinittest/a"
	"awesomeProject/myinittest/b"
)

func Test() {
	//c.Dis()
	//a:=a.A{"AAAAAA"}
	//b:=b.B{"BBBBB"}
	//a.DisPlay(b)

	aa := a.A{"AAAAAA"}
	b := b.B{"BBBBB"}
	aa.Dis(b)

}
