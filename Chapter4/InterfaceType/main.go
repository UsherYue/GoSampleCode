package main

import "fmt"

func main() {

	var var1 interface{} = 1
	var var2 interface{} = "hello,world"
	//map
	var var3 interface{} = map[string]string{"a": "1", "b": "2"}
	//匿名结构体
	var var4 interface{} = struct {
		x float64
		y float64
	}{x: 2, y: 2}
	var i = 1
	//匿名结构体指针
	var var5 interface{} = &i
	var var6 interface{} = &struct {
		x float64
		y float64
	}{x: 1, y: 2}
	fmt.Println(var1)
	fmt.Println(var2)
	fmt.Println(var3)
	fmt.Println(var4)
	fmt.Println(var5)
	fmt.Println(var6)
}
