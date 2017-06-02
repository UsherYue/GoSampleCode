package main

import (
	"fmt"
	"reflect"
)

var a2 = 1

//短变量不可以在全局使用
func main() {
	a3 := 1
	var a1 int = 1
	fmt.Println("a1=", a1, " type:", reflect.TypeOf(a1).Name())
	fmt.Println("a2=", a1, " type:", reflect.TypeOf(a2).Name())
	fmt.Println("a3=", a1, " type:", reflect.TypeOf(a3).Name())

	var b1 string = "hello"
	b2 := "hello"
	var b3 = "hello"
	fmt.Println("b1=", b1, " type:", reflect.TypeOf(b1).Name())
	fmt.Println("b2=", b1, " type:", reflect.TypeOf(b2).Name())
	fmt.Println("b3=", b1, " type:", reflect.TypeOf(b3).Name())

	var c1 = []int{1, 2, 3}
	var c2 []int = []int{1, 2, 3}
	c3 := []int{1, 2, 3}
	fmt.Println("c1=", c1, " type:", reflect.TypeOf(c1).Name())
	fmt.Println("c2=", c1, " type:", reflect.TypeOf(c2).Name())
	fmt.Println("c3=", c1, " type:", reflect.TypeOf(c3).Name())

}
