package myreftest

import (
	"fmt"
	"reflect"
)

//https://studygolang.com/articles/12348  学习资料   https://github.com/Jingbao1314/gonote
type Person struct {
	age int
}

type Student struct {
	Person
	Id   int
	Name string
}

func (s Student) GetId() (ids int, names string) {
	return s.Id, s.Name

}
func (s Student) getName() string {
	return s.Name

}

func MyRef() {
	stu := Student{Person{21}, 11, "xx"}
	key := reflect.TypeOf(stu)
	value := reflect.ValueOf(stu)
	fmt.Println(key)
	fmt.Println(value)
	for i := 0; i < key.NumField(); i++ {
		fmt.Print(key.Field(i), "---", key.Field(i).Offset, "-----")
		fmt.Println(value.Field(i))

	}
	//获得Student中可见的方法`
	fmt.Println(key.NumMethod())
	for j := 0; j < key.NumMethod(); j++ {
		fmt.Println(key.Method(j).Name)

	}
	//获取Student的所有方法
	key1 := reflect.TypeOf(&stu)
	for j := 0; j < key1.NumMethod(); j++ {
		fmt.Println(key1.Method(j).Name, "--------")

	}

	//
	stus := reflect.New(reflect.TypeOf(Student{}))
	//defer func() {
	//	err:=recover()
	//	if err!=nil{
	//		fmt.Println(err,"------------------")
	//		os.Exit(1)
	//	}
	//}()
	//stus.Elem().FieldByName("id").SetInt(11)
	//stus.Field(1).FieldByName("id").SetInt(11)
	stus.Elem().FieldByName("Name").SetString("xxx")
	fmt.Println("------------------", stus.Elem())
	strName := "Student"
	//反射
	if strName == "Student" {
		stuOne := reflect.New(reflect.TypeOf(Student{}))
		stuOne.Elem().FieldByName("Id").SetInt(11)
		stuOne.Elem().FieldByName("Name").SetString("陈润民")
		stuOneValue := reflect.ValueOf(stuOne)
		for ii := 0; ii < stuOneValue.NumField(); ii++ {
			fmt.Println(stuOneValue.Field(ii), "么么")
		}

	}

}
