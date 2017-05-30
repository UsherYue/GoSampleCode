package main

import "fmt"

func main() {
	var a1 int = 1
	var a2 int32
	var a3 float32
	var a4 complex64 = complex(1, 2.9)
	var a5 string
	var a6 byte
	var a7 [5]int //数组在golang中是按值传递而不是按地址传递
	var a8 map[string]string = map[string]string{"zhangsan": "18"}
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)
	fmt.Println(a5)
	fmt.Println(a6)
	fmt.Println(a7)
	fmt.Println(a8)
	fmt.Println("------------------")
	a2 = 100
	a5 = "golang"
	a6 = 254
	a7[0] = 100
	a8["lisi"] = "20"
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)
	fmt.Println(a5)
	fmt.Println(a6)
	fmt.Println(a7)
	fmt.Println(a8)
	fmt.Println("------------------")
	a1 = int(a2)
	a1 = int(a3)
	fmt.Println("------------------")
	var (
		b1 int    = 1
		b2 int    = 2
		b3 string = "hello"
	)
	fmt.Println(b1)
	fmt.Println(b2)
	fmt.Println(b3)
}
