package main

import "fmt"

func main() {

	var str interface{} = "hello,world!"
	var num interface{} = 1
	if v, ok := str.(string); ok {
		fmt.Println("str是string类型,str=", v)
	} else {
		fmt.Println("str不是string类型")
	}

	if v, ok := num.(string); ok {
		fmt.Println("num是string类型,num=", v)
	} else {
		fmt.Println("num不是string类型")
	}

}
