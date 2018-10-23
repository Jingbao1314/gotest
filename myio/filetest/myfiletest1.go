package filetest

import (
	"fmt"
	"os"
	"os/exec"
)

func MyReadFile() {
	myfile := "/home/jingbao/桌面/scala语法"
	in, err := os.Open(myfile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer in.Close()
	buff := make([]byte, 1024)
	for {
		flag, _ := in.Read(buff)
		//if 0==flag{
		//	break //如果读到个数为0,则读取完毕,跳出循环
		//}
		////从缓冲slice中写出数据,从slice下标0到n,通过os.Stdout写出到控制台
		//os.Stdout.Write(buff[:flag])
		if flag > 0 {
			os.Stdout.Write(buff[:flag])
		} else {
			break
		}

	}
	fmt.Println("---------------------------------------------------------------------------------------------------")
	userFile := "/home/jingbao/桌面/scala语法" //文件路径
	fin, err := os.Open(userFile)          //打开文件,返回File的内存地址
	defer fin.Close()                      //延迟关闭资源
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	buf := make([]byte, 1024) //创建一个初始容量为1024的slice,作为缓冲容器
	for {
		//循环读取文件数据到缓冲容器中,返回读取到的个数
		n, _ := fin.Read(buf)

		if 0 == n {
			break //如果读到个数为0,则读取完毕,跳出循环
		}
		//从缓冲slice中写出数据,从slice下标0到n,通过os.Stdout写出到控制台
		os.Stdout.Write(buf[:n])
	}
}

func MyWriteTest() {
	userFile := "/home/jingbao/桌面/scala语法"
	out, err := os.Create(userFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer out.Close()
	out.WriteString("jingbao")
	out.Write([]byte{'j', 'i', 'n', 'g'})

}

func MyLoader() {
	usePath := "/home/jingbao/gocode/src/awesomeProject/myio/filetest/xxx.go"
	in, err := os.Create(usePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer in.Close()
	files := "package filetest \n" + "import (\n   \"fmt\" \n) \n" + "type Person struct{\n   Age int \n}\n \n \n" + "func (this Person)Say(){ \n" + "   fmt.Println(7721) \n}"
	files = files + "\n \n" + "func FuncTest(){ \n" + "   person:=Person{11} \n   person.Say() \n}"
	in.WriteString(files)

	cmd := exec.Command("go build", "/home/jingbao/gocode/src/awesomeProject/myio/filetest/xxx.go")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))

}
