package myexception

import "fmt"

//func GO() {
//	fmt.Println("我是GO，现在没有发生异常，我是正常执行的。")
//}
//
//func PHP() {
//	// panic一般会导致程序挂掉（除非recover）  然后Go运行时会打印出调用栈
//	//但是，关键的一点是，即使函数执行的时候panic了，函数不往下走了，运行时并不是立刻向上传递panic，而是到defer那，等defer的东西都跑完了，panic再向上传递。所以这时候 defer 有点类似 try-catch-finally 中的 finally。panic就是这么简单。抛出个真正意义上的异常。
//	panic("我是PHP,我要抛出一个异常了，等下defer会通过recover捕获这个异常，然后正常处理，使后续程序正常运行。")
//	fmt.Println("我是PHP里panic后面要打印出的内容。")
//}
//
//func PYTHON() {
//	fmt.Println("我是PYTHON，没有defer来recover捕获panic的异常，我是不会被正常执行的。")
//}

func GO() {
	fmt.Println("我是GO，现在没有发生异常，我是正常执行的。")
}

func PHP() {
	// 必须要先声明defer，否则不能捕获到panic异常,也就是说要先注册函数，后面有异常了，才可以调用
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("终于捕获到了panic产生的异常：", err) // 这里的err其实就是panic传入的内容
			fmt.Println("我是defer里的匿名函数，我捕获到panic的异常了，我要recover，恢复过来了。")
		}
	}() //注意这个()就是调用该匿名函数的，不写会报expression in defer must be function call

	// panic一般会导致程序挂掉（除非recover）  然后Go运行时会打印出调用栈
	//但是，关键的一点是，即使函数执行的时候panic了，函数不往下走了，运行时并不是立刻向上传递panic，而是到defer那，等defer的东西都跑完了，panic再向上传递。所以这时候 defer 有点类似 try-catch-finally 中的 finally。panic就是这么简单。抛出个真正意义上的异常。
	panic("我是PHP,我要抛出一个异常了，等下defer会通过recover捕获这个异常，捕获到我时，在PHP里是不会输出的，会在defer里被捕获输出，然后正常处理，使后续程序正常运行。但是注意的是，在PHP函数里，排在panic后面的代码也不会执行的。")
	fmt.Println("我是PHP里panic后面要打印出的内容。但是我是永远也打印不出来了。因为逻辑并不会恢复到panic那个点去，函数还是会在defer之后返回，也就是说执行到defer后，程序直接返回到main()里，接下来开始执行PYTHON()")
}

func PYTHON() {
	fmt.Println("我是PYTHON，没有defer来recover捕获panic的异常，我是不会被正常执行的。")
}

func Test() {
	GO()
	PHP()
	PYTHON()

}
