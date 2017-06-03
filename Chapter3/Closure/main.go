package main

import "fmt"

func TestCallback(callback func(string)) {
	callback("hello,callback!")
}

//这里就是将匿名函数当做函数的返回值来使用
func GetCallback() func(string) {
	return func(str string) {
		fmt.Println(str)
	}
}

//费波纳西额数列的实现
func Fibonacci() func() int {
	i, j := 0, 0
	return func() int {
		if i == 0 && j == 0 {
			i = 1
			return 1
		} else if i == 1 && j == 0 {
			j = 1
			return 1
		} else {
			tmp := i
			i += j
			j = tmp
			return i
		}
	}
}

func main() {
	//定义一个短变量来保存匿名函数
	//	callback := func(str string) {
	//		fmt.Println(str)
	//	}
	callback := GetCallback()
	//这也就是说明  匿名函数可以当做变量 当做函数参数 也可以当做函数返回值来使用
	TestCallback(callback)

	fibonacci := Fibonacci()
	for i := 0; i < 100; i++ {
		fmt.Printf("%d ", fibonacci())
	}

}
