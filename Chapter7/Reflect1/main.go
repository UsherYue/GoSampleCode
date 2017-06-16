/*
获取一个结构体对象的方法和属性列表
*/
package main

import (
	"fmt"
	"os"
	"reflect"
)

type DemoStruct struct {
	A, B, C int
	D, E, F string
	H, I, J []string
}

func (this DemoStruct) TestFunc1() {
	fmt.Println("Test Func1....")
}
func (this DemoStruct) TestFunc2() {
	fmt.Println("Test Func2....")
}
func (this DemoStruct) TestFunc3() {
	fmt.Println("Test Func3....")
}

func main() {
	item := &DemoStruct{}
	tp := reflect.TypeOf(*item)
	fmt.Println("Kind:", tp.Kind().String())
	fmt.Println("TypeName::", tp.Name())
	fmt.Fprintln(os.Stdout, "字段列表:")
	for i := 0; i < (tp).NumField(); i++ {
		fmt.Println((tp).Field(i).Name)
	}
	fmt.Fprintln(os.Stdout, "方法列表:")
	for i := 0; i < tp.NumMethod(); i++ {
		fmt.Println(tp.Method(i).Name)
	}

}
