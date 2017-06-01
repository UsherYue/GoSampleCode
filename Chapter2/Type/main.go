package main

import (
	"fmt"
	"math"
	"os"
)

//关于布尔类型与关系逻辑运算的测试
func testBoolean() {
	//> < != == >= <=
	var b1 bool = (1 > 2)  //false
	var b2 bool = (1 < 2)  //true
	var b3 bool = (3 == 3) //frue
	var b4 bool = (1 <= 5) //true
	var b5 bool = (1 >= 5) //false
	var b6 bool = (1 != 5) //true
	fmt.Fprintf(os.Stdout, "a1=%t,a2=%t,a3=%t,a4=%t,a5=%t,a6=%t\n", b1, b2, b3, b4, b5, b6)
	//下面我们使用变量 以及字面常量进行关系运算
	var name string = "usher"
	var b7 bool = (name == "usher")
	fmt.Println("b7=", b7)
	var (
		b8 int = 10
		b9 int = 100
	)
	fmt.Println("b8==b9:", b8 == b9)
	// &&  || !
	//A&&B  A||B !A
	fmt.Println("b1&&b2=", b1 && b2)
	fmt.Println("b1||b2=", b1 || b2)
	fmt.Println("!b1=", !b1)
	//复杂一点的逻辑运算
	var b10 = b1 && b2 && b3 //false
	fmt.Println("b1 && b2 && b3=", b10)
	var b11 = b1 || b2 && b3 //true
	fmt.Println("b1 || b2 && b3=", b11)
}

func main() {
	testBoolean()
}

//整型测试函数
func testInterger() {
	//有符号   int8 int16 int32
	var a1 int8 = 127
	var a2 int16 = -10000
	var a3 int32 = 5000
	//无符号类型  byte uint8 uint16
	var b1 byte = 100
	var b2 uint8 = 255
	var b3 uint16 = 65535
	fmt.Fprintf(os.Stdout, "a1=%d,a2=%d,a3=%d\n", a1, a2, a3)
	fmt.Fprintf(os.Stdout, "b1=%d,b2=%d,b3=%d\n", b1, b2, b3)
	// 我们如何将a1 赋值给a2
	// 类型转换 a->b   b(a)
	a2 = int16(a1)
	b3 = uint16(b1)
	fmt.Fprintf(os.Stdout, "a1=%d,a2=%d,a3=%d\n", a1, a2, a3)
	fmt.Fprintf(os.Stdout, "b1=%d,b2=%d,b3=%d\n", b1, b2, b3)

	//类型转换不是无限制的，是有前提和条件的  我们不能将string类型和整数类型相互转换 、也不能将数组类型和整数类型相互转换 这是常识问题
	///var str string = "hello"
	//fmt.Println(int(str))
	//var arr []int = []int{1, 2, 3, 4}
	//fmt.Println(int(arr))
}

//浮点数测试
func testFloat() {
	var f1 float64 = 1.1
	var f2 float32 = 1.2
	fmt.Fprintf(os.Stdout, "f1=%.2f,f2=%.2f\n", f1, f2)
	//我们刚刚说了 浮点数它是一种不精确的表示方式 所以我们不能直接  浮点数1==浮点数2
	var (
		c1 float32 = 1.11
		c2 float32 = 1.11000000001
	)
	fmt.Println("c1==c2:", c1 == c2)
	//我们知道浮点数的比较是不可能精确地所以说我们假设如果两个浮点数的差值等于0 或者 我们自定义的一个值那么我们就认为他们是相等的
	var equalValue float64 = 0.00000001
	if math.Dim(math.Max(float64(c1), float64(c2)), math.Min(float64(c1), float64(c2))) < equalValue {
		fmt.Println("c1==c2")
	} else {
		fmt.Println("c1!=c2")
	}
}
