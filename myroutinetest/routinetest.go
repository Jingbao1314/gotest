package myroutinetest

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello world goroutine")
}
func Test() {
	go hello()
	time.Sleep(1 * time.Second)
	fmt.Println("main function")
}
