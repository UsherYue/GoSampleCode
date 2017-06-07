package main

import (
	"fmt"
	"time"
)

func Work(i int) {
	fmt.Println("Do Work Index:", i)
}

func main() {
	//由于现在go Work(i)调用时并发执行 所以这里的i肯定是0~99之间的乱序输出
	for i := 0; i < 100; i++ {
		go Work(i)
	}
	//主协程休眠10毫秒 防止子协程没有执行完毕 主协程退出
	time.Sleep(time.Millisecond * 10)
}
