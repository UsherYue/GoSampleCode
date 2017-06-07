package main

import (
	"fmt"
	//	"time"
)

type TaskData struct {
	Cmd  string
	Data string
}

func AsyncTask(data interface{}) {
	task := data.(*TaskData)
	if task.Cmd == "End" {
		fmt.Println("------------收到退出指令", task.Data, "-------------")
	} else {
		//过多的fmt.Println会导致系统调用堵塞 这里注意
		fmt.Println("Cmd:", task.Cmd, " Data:", task.Data)
	}
}

func main() {
	queue := NewQueue(10, AsyncTask)
	//投递任务
	for i := 0; i < 1000; i++ {
		//time.Sleep(100)
		if i == 999 {
			queue.Post(&TaskData{"End", fmt.Sprintf("This is a End Message,Index:%d", i)})
		} else {
			queue.Post(&TaskData{"Send", fmt.Sprintf("This is a Send Message,Index:%d", i)})
		}
	}
	//5s后停止
	queue.StopAfter(5)
	//启动队列
	queue.Start()
}
