package myinterface

import (
	"fmt"
	"reflect"
	"strconv"
)

//https://studygolang.com/articles/15278
//https://studygolang.com/articles/5771  接收者
type t1 int
type t2 int

//func( t *t2) String() string { return "ptr" }
//func( t t1) String() string { return "val" }
func (t *t1) String() string { return "ptr" }
func (t t2) String() string  { return "val" }
func Test() {
	fmt.Println("---------------------")
	var a t1
	var b t2
	a = 5
	fmt.Println(a, b)
}

type Person struct {
	Age int
}

func (p Person) Say() {
	fmt.Println("Say")

}

func (p *Person) Says() {
	fmt.Println("Says")

}

func (p Person) SayOk() {
	fmt.Println("Ok")
}

func (p *Person) SayYes() {
	fmt.Println("Yes")

}

type Stu struct {
	Person
	Name string
}

func (p Stu) SayOk() {
	fmt.Println("Stu Ok")
}

func (p *Stu) SayYes() {
	fmt.Println("Stu Yes")
}

func Test1() {
	//p:=Person{11}
	//s:=Stu{p,"xx"}
	////s:=Stu{Person{11},"qqqq"}
	//stype:=reflect.TypeOf(s)
	//smethod:=stype.NumMethod()
	//for i:=0;i<smethod ;i++  {
	//	fmt.Println(stype.Method(i).Name)
	//}
	p := Person{11}
	s := Stu{p, "xx"}
	//s:=Stu{Person{11},"qqqq"}
	stype := reflect.TypeOf(&s)
	smethod := stype.NumMethod()
	for i := 0; i < smethod; i++ {
		fmt.Println(stype.Method(i).Name)
	}

}

type Father1 interface {
	Say()
}

type Son1 struct {
	Id int
}

func (this Son1) Say() {
	fmt.Println("Son1 say")

}

type Father2 interface {
	Play()
}

type Son2 struct {
	Id int
}

func (this Son2) Say() {
	fmt.Println("Son2 say")

}
func (this Son2) Play() {
	fmt.Println("Son2 Play")

}

func Test2() {
	//s1:=Son1{7721}
	//s2:=Son2{8888}
	////f1:=unsafe.Pointer(&s1)
	////f2:=unsafe.Pointer(&s2)
	//fmt.Println(reflect.TypeOf(s1).(Son1).Id)
}

type Student interface {
	ResoveQues(call MyCall)
}

type MyCall interface {
	tellAnswer(answer int)
}

type Teacher struct {
	Stus Student
}

func (this Teacher) tellAnswer(answer int) {
	fmt.Println("答案是:" + strconv.Itoa(answer))
}
func (this Teacher) askAnswer() {
	this.Stus.ResoveQues(this)
}

type Rookie struct {
}

func (this Rookie) ResoveQues(call MyCall) {
	call.tellAnswer(7721)

}

func Test3() { //回调
	var t = Teacher{}
	var s = Rookie{}
	t.Stus = s
	t.askAnswer()

}
