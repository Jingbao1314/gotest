package myroutinetest

import (
	"fmt"
	"strconv"
	"time"
)

func hello() {
	fmt.Println("Hello world goroutine")
}
func Test() {
	go hello()
	time.Sleep(1 * time.Second)
	fmt.Println("main function")
	fmt.Println("----------------------------------------------------------------1")
	c := make(chan int)
	go chan1(&c)
	v, _ := <-c
	fmt.Println(v)
	fmt.Println("----------------------------------------------------------------2")
	c2 := make(chan int, 2)
	go chan2(&c2)
	c2 <- 10
	c2 <- 100
	//time.Sleep(2*time.Second)
	v1, _ := <-c2
	fmt.Println("res is :" + strconv.Itoa(v1))
	fmt.Println("----------------------------------------------------------------3")
}

func chan1(c *chan int) {
	*c <- 10
}

func chan2(c *chan int) {
	time.Sleep(time.Second * 1)
	for i := 0; i < 2; i++ {
		v, _ := <-*c
		fmt.Println("Read chan : " + strconv.Itoa(v))
	}
	*c <- 7721
}
