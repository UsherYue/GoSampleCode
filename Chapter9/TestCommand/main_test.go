package main

import (
	"fmt"
	"testing"
)

//函数必须大写Test开头 ，并且有一个 testing.T的指针类型的参数
//	testing.B
//  testing.TB
func TestFunc(t *testing.T) {
	fmt.Println("go test is called!")

}
