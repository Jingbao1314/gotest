package main

import "awesomeProject/myinterface"

//https://studygolang.com/articles/13031 编译相关
//type mystring struct {
//	Lens int16
//	Pos int32
//}
//
//func xx() *int32 {
//	var a int32=6
//	var b *int32 =&a
//	return b
//}
//
//func xxx()  (int8,int8){
//	return 1,2
//
//}
/**
pkg 里放编译生成的连接库/

bin 编译生成的可执行程序

go run  边build边执行
数组  切片

*/

func main() {
	//var xx mystring
	//fmt.Println(unsafe.Sizeof(xx))
	//fmt.Println(unsafe.Alignof(xx))
	//fmt.Println(unsafe.Offsetof(xx.Lens))
	//println(unsafe.Sizeof(xx))//大小
	//println(unsafe.Alignof(xx))//对齐规则
	//println(unsafe.Offsetof(xx.Lens))//属性偏移位置
	//println(unsafe.Offsetof(xx.Pos))

	//var as=[...]int32{1,2,3}
	//var as []int32=make([]int32,6)//切片
	//as[0]=1
	//as=append(as,3)
	//fmt.Println(cap(as), len(as))
	//as[6]=1
	//as[2]=1
	//as[5]=1
	//as=append(as,3)
	//as=append(as,3)
	//as=append(as,3)
	//as=append(as,3)
	//as=append(as,3)
	//as=append(as,3)
	//fmt.Println(cap(as), len(as))
	//
	//var as1  [3]int32
	//as1[0]=0
	//var as2 [3]int32
	//as2[0]=0
	//fmt.Println(reflect.TypeOf(as),reflect.TypeOf(as1))
	//fmt.Println(reflect.TypeOf(as).Kind(),reflect.TypeOf(as2).Kind(),reflect.TypeOf(as1).Kind())

	//var str [3]mystring
	//var x1  mystring
	//str[1]=x1
	//var str1 [3]mystring
	//var x2  mystring
	//str1[1]=x2
	//
	//fmt.Println(reflect.TypeOf(str),reflect.TypeOf(str1))
	//fmt.Println(reflect.TypeOf(str).Kind(),reflect.TypeOf(str1).Kind())

	//var a *int16=new(int16)
	//var b  int16=0x0102
	//a=&b
	//var c uintptr = uintptr(unsafe.Pointer(a))
	//var d  =c+1
	//var c1  = (*int8)(unsafe.Pointer(c))
	//var d1  = (*int8)(unsafe.Pointer(d))//void *
	//
	//fmt.Println(*c1,*d1)

	//var a *int16=new(int16)
	//b:=10
	//a=&b
	//fmt.Println(*a)

	//var arr=[...]int8{1,2,3,4,5}//切片运算
	//var sli  = arr[2:4]
	//fmt.Println(sli)
	//fmt.Println(reflect.TypeOf(arr))

	//var a=[...]int{7:4,1}
	//var b  = &a
	//fmt.Println(b[7])
	//for k,v:= range a{
	//	fmt.Println(k,v)
	//}
	//fmt.Println(reflect.TypeOf(a).Kind())

	//-l拒绝内连  -N防止优化
	//-gcflags "-m -l"
	//go tool objdump -s "main\.main" aaa
	// go build -gcflags "-l -N" test.go
	//反汇编
	//go build test.go   go tool objdump -s "main\.main" test
	//var aa  = 0x1111
	//var bb  =0x2222
	//var c  =aa+bb
	//fmt.Println(c)

	//if a:=6;a>5{
	//	fmt.Println(">5")
	//}else {
	//	fmt.Println("xxxxxxx")
	//}

	//switch a:=6; {
	//case a>5:
	//	fmt.Println(">5")
	//	//fallthrough
	//case a>10:
	//	fmt.Println(">10")
	//default:
	//	fmt.Println("xxxxxx")
	//}

	//var arr   = make([]int8,5)
	//for k,v:=range arr{
	//	fmt.Println(k,v)
	//}

	//string  转  int    strconv.Atoi(int)
	//var  a=1
	//var c ="a="+strconv.Itoa(a)
	//fmt.Println(c)

	//a:="1"
	//c, _ :=strconv.Atoi(a)
	//fmt.Println(c+1)

	//多返回值
	//var a,b=xxx()
	//fmt.Println(a,b)

	//const (
	//	a = iota   //0
	//	b          //1
	//	c          //2
	//	d = "ha"   //独立值，iota += 1
	//	e          //"ha"   iota += 1
	//	f = 100    //iota +=1
	//	g          //100  iota +=1
	//	h = iota   //7,恢复计数
	//	i          //8
	//)
	//fmt.Println(a,b,c,d,e,f,g,h,i)
	//mylang.Strtest()
	//var a=[...]int{2,7,11,15,98}
	//mylang.Slicetest()

	//函数test
	//var a = []interface{}{123, "abc"}
	//
	//mylang.Print(a...) // 123 abc
	//mylang.Print(a)    // [123 abc]
	//c:=7721
	//
	//var b= func() {
	//	fmt.Println("匿名函数:"+strconv.Itoa(c))
	//}
	//b()
	//
	//var d []int=[]int{1,2,3}
	//fmt.Println(d)

	//fmt.Println(mylang.Inc())

	//mylang.DeferTest()

	//var a []int=make([]int,5)
	//a[0]=1
	//a[1]=2
	//a[2]=7
	//a[3]=15
	//a[4]=19
	//var b int=9
	//myleetcode.TwoSum(a,b)

	//myleetcode.Test()
	//mylang.Slicetest()
	//fmt.Println(mylang.Inc())
	//filetest.MyWriteTest()
	//filetest.MyReadFile()
	//myreftest.MyRef()
	//myref.Myref()
	//car:=mylang.BigCar{}
	//car.Start()
	//
	//files:="package filetest \n"+"import (\n   fmt \n) \n"+"func TestFile(){ \n"+"   fmt.Println(7721) \n}"
	//fmt.Print(files)
	//filetest.MyLoader()
	//filetest.-FuncTest()
	//myroutinetest.Test()
	//f:=filetest.Say
	//f()

	//fmt.Println(myleetcode.StrStr("acc","cc"))
	//x:=make([][]int,3)
	//x[0]=[]int{0,1,2}
	//fmt.Println(x)
	//myleetcode.Convert("qwertyuiop",4)
	//buff:=make([][]byte,2)
	//buff[0]=[]byte{1,'.',1}
	//buff[1]=[]byte{1,2,3}
	//myleetcode.IsValidSudoku(buff)

	//mylang.ResTest()
	myinterface.Test1()
	//type PipeData struct {
	//	value   int
	//	handler func(a int) int
	//}
	//
	////实现不了的部分⬇
	//
	//var a PipeData
	//
	////fmt.Print(a.handler(2))
	//
	////实现不了的部分⬆
	//var myhandler = func(a int) int {
	//	return a*2
	//}
	//
	//a.handler = myhandler
	//fmt.Println(a.handler(2))
}
