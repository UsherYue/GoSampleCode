package main

import (
	"fmt"
	"unsafe"
)

type Base struct {
	i int
}

func (this *Base) MethodA() {
	fmt.Println("Base::Method A is Called!")
}

type PSub struct {
	*Base
	i int
}

type Sub struct {
	Base
	i int
}

func (this *PSub) MethodA() {
	fmt.Println("PSub::Method A is Called!")
}

func (this *Sub) MethodA() {
	fmt.Println("Sub::Method A is Called!")
}
func main() {
	//匿名组合
	fmt.Println("--------------匿名组合--------------")
	//我们定义一个Sub类型的对象 并且将指针传递给pSub  i初始化为1
	pSub := &Sub{i: 100}
	//我们调用的MethodA实际上是Sub中的MethodA这个方法
	pSub.MethodA()
	//我们可以调用到父结构体的方法或者属性
	pSub.Base.MethodA()
	fmt.Println("Base::i=", pSub.Base.i)
	fmt.Println("Sub::i=", pSub.i)
	//指针匿名组合继承
	fmt.Println("--------------通过指针匿名组合--------------")
	pPSub := &PSub{Base: &Base{}}
	pPSub.MethodA()
	pPSub.Base.MethodA()
	fmt.Println("--------------父结构体指针和子结构体指针相互转换--------------")
	var pBase *Base
	//将子结构体指针转换成父结构体指针
	//unsafe.Pointer 实际上本质上是一个int类型的指针
	pBase = (*Base)(unsafe.Pointer(pSub))
	pBase.MethodA()
	//将父结构体指针转换成子结构体指针
	(*Sub)(unsafe.Pointer(pBase)).MethodA()
	(*Sub)(unsafe.Pointer(pBase)).Base.MethodA()

}
