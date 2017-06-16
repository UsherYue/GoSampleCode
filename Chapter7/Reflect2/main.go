/*
动态调用方法
*/
package main

import (
	"fmt"
	"reflect"
)

type DemoStruct struct {
}

func (this DemoStruct) TestFunc1(a int) int {
	fmt.Println("Test Func1....")
	return a
}
func (this DemoStruct) TestFunc2() {
	fmt.Println("Test Func2....")
}
func (this DemoStruct) TestFunc3() {
	fmt.Println("Test Func3....")
}

func main() {
	item := DemoStruct{}
	// 获取对象的 reflect.Value
	val := reflect.ValueOf(item)
	a := val.MethodByName("TestFunc1").Call([]reflect.Value{reflect.ValueOf(2)})
	val.MethodByName("TestFunc2").Call(nil)
	val.MethodByName("TestFunc3").Call(nil)
	fmt.Println("TestFunc1 Return:", a[0].String())
	//定义一个变量 通过反射来动态设置
	b := 100
	reflect.ValueOf(&b).Elem().SetInt(1000)
	fmt.Println("b=", b)
}
