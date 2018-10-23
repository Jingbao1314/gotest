package mylang

import (
	"fmt"
)

// 具名函数
func Add(a, b int) int {
	return a + b
}

//// 匿名函数
//var Add = func(a, b int) int {
//	return a+b
//}

// 多个输入参数和多个返回值
func Swap(a int, b int) (int, int) {
	return b, a
}

// 可变数量的参数
// more 对应 []int 切片类型
func Sum(a int, more ...int) int {
	for _, v := range more {
		a += v
	}
	return a
}

//带名字 的返回值
func Find(m map[int]int, key int) (value int, ok bool) {
	value, ok = m[key]
	return
}

//闭包  defer:延迟
func Inc() (v int) {
	defer func() { v++ }()
	return 42
}

func Test() int {
	defer fmt.Println("3")
	defer fmt.Println("2")
	defer fmt.Println("1")
	return 10
}

func DeferTest() {
	//for i:=0;i<3;i++{
	//	defer func() {fmt.Println(i)}()
	//}
	defer func() { fmt.Println(3) }()
	defer func() { fmt.Println(2) }()
	defer func() { fmt.Println(1) }()
	//for i:=0;i<3;i++{
	//	i:=i
	//	defer func() {fmt.Println(i)}()
	//}
	//for i:=0;i<3;i++{
	//	defer func(i int) {fmt.Println(i)}(i)
	//}
}

type Engine interface {
	Start()
}

type Car struct {
	Engine
}

type BigCar struct {
	Car
}

func (this *BigCar) Start() {
	fmt.Println("BigCar->Start()")
}

func (this *Car) Start() {
	fmt.Println("Car->Start()")
}

//接收者
type test struct {
	name string
}

func (t test) TestValue() {
	fmt.Printf("%s\n", t.name)
}

func (t *test) TestPointer() {
	fmt.Printf("%s\n", t.name)
}

//https://github.com/qyuhen/book
func ResTest() {
	t := test{"wang"}
	m := t.TestValue //这里发生了复制,不受后面修改的影响
	t.name = "Li"
	m1 := (*test).TestPointer
	m1(&t) //Li
	m()    //wang
	m = t.TestValue
	m() //Li
}
