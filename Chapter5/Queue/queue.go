package main

import (
	"time"
)

type WorkFunc func(interface{})

type Queue struct {
	ch chan interface{}
	//进程控制退出标志
	exitFlag chan bool
	onWork   WorkFunc
}

func (this *Queue) StopAfter(seconds int) {
	go func() {
		time.Sleep(time.Second * time.Duration(seconds))
		this.exitFlag <- true
	}()
}

func (this *Queue) Block() {
	<-this.exitFlag
}

func (this *Queue) Post(data interface{}) {
	go func() {
		this.ch <- data
	}()
}

//启动队列
func (this *Queue) Start() {
	go func() {
		for {
			select {
			case data := <-this.ch:
				{
					if this.onWork != nil {
						go this.onWork(data)
					}
				}
			default:
				{

				}
			}
		}
	}()
	this.Block()
}

func NewQueue(size int, onWork WorkFunc) *Queue {
	return &Queue{ch: make(chan interface{}, size), exitFlag: make(chan bool), onWork: onWork}
}
