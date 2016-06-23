package main

//https://golang.org/doc/gdb
import (
	"fmt"
)

// 因为golang程序在编译的时候会生成DWARF3 这样一个版本的调试信息  > 7.1的gdb都可以进行调试。
func main() {
	mp := map[string]int{"a": 1, "b": 2}
	fmt.Println(mp)
	fmt.Println("请输入数字:")
	var i int
	fmt.Scanf("%d", &i)
	fmt.Printf("输入的数字是:%d\n", i)
	for j := 0; j < 100; j++ {
		i += j
		fmt.Println("i+=j:", i)
	}
}
