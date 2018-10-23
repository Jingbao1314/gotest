package xx

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func Xxx(a int32) {
	var str string = "xxxx"

	fmt.Println(len([]rune(str)))
	var mylist [5]int32
	mylist[0] = 1
	fmt.Println(reflect.TypeOf(mylist))
	fmt.Println(reflect.TypeOf(mylist).Kind()) //数据结构  种类
}

type User struct {
	Count int
	Name  string
}

func Atm() {
	var user User
	user.Count = 0
	user.Name = "xxx先生"
	for {
		fmt.Println("请选择业务,1:存,2:取,3:查")
		var num int
		fmt.Scanln(&num)
		if num == 1 {
			goto saves
		} else if num == 2 {
			goto take
		} else if num == 3 {
			goto selects
		} else {
			fmt.Println("再见")
			os.Exit(0)
		}
	saves:
		{
			fmt.Println(user.Name + ",您的余额为:" + strconv.Itoa(user.Count) + ",确认无误后请输入您要存数的金额")
			var money int
			fmt.Scanln(&money)
			user.Count = user.Count + money
			fmt.Println("存储成功,您的余额为:" + strconv.Itoa(user.Count) + ",确认无误后请选择业务,1:存,2:取,3:查")
			var num int
			fmt.Scanln(&num)
			if num == 1 {
				goto saves
			} else if num == 2 {
				goto take
			} else if num == 3 {
				goto selects
			} else {
				fmt.Println("再见")
				os.Exit(0)
			}
		}
	take:
		{
			fmt.Println(user.Name + "您的余额为:" + strconv.Itoa(user.Count) + ",确认无误后请输入您要取款的金额")
			var money int
			fmt.Scanln(&money)
			var temp = user.Count - money
			if temp < 0 {
				fmt.Println("余额不足,您可以输入:1.进行充值,2.查看余额,3.退出")
				var i int
				fmt.Scan(&i)
				if i == 1 {
					goto saves
				} else if i == 2 {
					goto selects
				} else {
					fmt.Println("再见")
					os.Exit(0)
				}
			} else {
				user.Count = temp
			}
			fmt.Println("取款成功,您的余额为:" + strconv.Itoa(user.Count) + ",确认无误后请选择业务,1:存,2:取,3:查")
			var num int
			fmt.Scanln(&num)
			if num == 1 {
				goto saves
			} else if num == 2 {
				goto take
			} else if num == 3 {
				goto selects
			} else {
				fmt.Println("再见")
				os.Exit(0)
			}

		}
	selects:
		{
			fmt.Println(user.Name + "您的余额为:" + strconv.Itoa(user.Count) + ",确认无误后请选择业务,1:存,2:取,3:查")
			var num int
			fmt.Scanln(&num)
			if num == 1 {
				goto saves
			} else if num == 2 {
				goto take
			} else if num == 3 {
				goto selects
			} else {
				fmt.Println("再见")
				os.Exit(0)
			}
		}

	}

}
