package main

import (
	"fmt"
	"sync"
	"time"
)

//全局唯一锁
var onceLock = &sync.Once{}

func onceFunc() {
	fmt.Println("这函数只会执行一次")
}

//同步锁
//var mutex = &sync.Mutex{}

//读写锁
var mutex = &sync.RWMutex{}
var i = 0

func addInt() {
	mutex.Lock()
	i++
	fmt.Println("i:=", i)
	mutex.Unlock()
}

func main() {

	//尽管我们调用了多次 实际上他只执行了一次
	for i := 0; i < 100; i++ {
		onceLock.Do(onceFunc)
	}
	for i := 0; i < 100; i++ {
		go addInt()
	}
	time.Sleep(2 * time.Second)
}
