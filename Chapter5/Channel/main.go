package main

import (
	"fmt"
)

// 数据传输通道  次数可以用缓冲通道 也可以用同步通道
var ch chan int = make(chan int, 100)

//控制进程退出
var sw chan bool = make(chan bool)

func Recv(rdch <-chan int) {
	for {
		//接收数据
		i := <-rdch
		fmt.Println("Recv:", i)
		if i == 9 {
			sw <- true
			break
		}
	}
}

func main() {
	//启动数据接收goroutine
	go Recv(ch)
	for i := 0; i < 10; i++ {
		//写入数据到channel中
		ch <- i
	}
	<-sw
	//关闭通道
	close(sw)
	close(ch)
}
