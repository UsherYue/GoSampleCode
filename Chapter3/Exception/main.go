package main

import "fmt"

//defer 语句只能接函数调用 并且多个defer语句是按照后进先出的顺序进行调用的
func TestDefer() {
	var a = 1
	fmt.Println("a=", a)
	defer func() {
		a *= 100
		fmt.Println("a=", a)
	}()
	defer func() {
		a *= 5
		fmt.Println("a=", a)
	}()
}

func testPanic() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("没有捕获到异常")
		}
	}()
	panic("this is a panic")

}

func main() {
	//TestDefer()
	testPanic()
}
