package mySync

import (
	"fmt"
	"sync"
	"time"
)

func MyTicker() {
	done := make(chan bool)
	ticker := time.NewTicker(time.Second * 5)
	go func() {
		select {
		case m := <-ticker.C:
			fmt.Println(m)
		case <-time.After(3 * time.Second):
			fmt.Println("timed out")
		}

		//total := 0
		//for t := range ticker.C {
		//	fmt.Println("Tick at:", t)
		//	total += 1
		//	if total > 10 {
		//		break
		//	}
		//}
		done <- true
	}()
	<-done
	ticker.Stop()
	//fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
}

func MyWaitGroup() {
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		//time.Sleep(1*time.Second)
		fmt.Println("二号选手准备就绪")
		wg.Done()
	}()
	go func() {
		fmt.Println("一号选手准备就绪")
		wg.Done()
	}()
	go func() {
		//time.Sleep(1*time.Second)
		fmt.Println("三号选手准备就绪")
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("起跑")
}
