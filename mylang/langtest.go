package mylang

import (
	"fmt"
	"reflect"
)

func Strtest() {
	a := "jingbao"
	var b = [...]byte{'j', 'i', 'n', 'g', 'b', 'a', 'o'}
	fmt.Println(a[1:4], b[:4])
	fmt.Println([]byte(a))
	fmt.Println(string([]rune{19990, 30028}))
	//var x int32 = 30028
	//fmt.Println(string(x))

}

func Arraytest() {
	//var a = [...]int32{1,2,3}
	//var b =[...]int32{2:5,3:4}
	//var c  =[...]int32{2:5,3,4}

}

func Slicetest() {
	//var a []int32=make([]int32,5)
	var a []int32 = []int32{1, 2, 3, 4, 5}
	//var b=a[2:4]
	//fmt.Println(b, len(b), cap(b))
	var c []int32 = make([]int32, 0, 5)

	for _, v := range c {
		fmt.Println(v)
	}
	fmt.Println(len(c), cap(c))
	fmt.Println(reflect.TypeOf(a))
}

func Maptest() {

}

func Print(a ...interface{}) {
	fmt.Println(a...)
}
