package main

import (
	"fmt"
	"time"
)

var (
	ch1 = make(chan int) //同步通道
	ch2 = make(chan int) // 同步通道
)

func main() {
	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			ch2 <- i
		}
	}()
	go func() {
		for {
			select {
			case i := <-ch1:
				{
					fmt.Println("ch1:", i)
				}
			case i := <-ch2:
				{
					fmt.Println("ch2:", i)
				}
			}
		}
	}()
	//主协程序暂停2s
	time.Sleep(time.Second * 6)
}
