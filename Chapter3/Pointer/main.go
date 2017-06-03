package main

import "fmt"

func AddByPointer(pi *int, val int) {
	*pi += val
}

func AddByValue(i int, val int) {
	i += val
}

func main() {

	////关于指针的创建
	p1 := new(int) //默*p1认就是0
	i := 0
	p2 := &i
	str := "zhangsan"
	p3 := &str
	fmt.Println("p1=", p1, " p2=", p2, " p3=", p3)
	fmt.Println("*p1=", *p1, " *p2=", *p2, " *p3=", *p3)
	///指针修改对应内存地址存储的内容
	*p3 = "lisi"
	fmt.Println("str=", str)
	*p1 = 1000
	fmt.Println("*p1=", *p1)
	*p2 = 200
	fmt.Println("i=", i)
	//指针当做函数参数传递
	i = 0
	AddByValue(i, 10)
	fmt.Println(i)
	AddByPointer(&i, 10)
	fmt.Println(i)

}
