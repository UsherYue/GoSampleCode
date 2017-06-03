package main

import "fmt"

func main() {

	fmt.Println("-------------go语言中数组的创建-------------------")
	//首先我们练习一下数组的创建
	var arr1 [6]int = [6]int{1, 2, 3, 4, 5, 6}
	//中括号中的...会被编译器自动替换成后面初始化列表中元素的个数
	var arr2 = [...]int{1, 2, 3, 4, 5}
	arr3 := [...]int{1, 2, 3, 4, 5}
	arr4 := [...]string{"zhangsan", "lisi", "wangwu"}
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)
	fmt.Println("-------------go语言中切片的创建-------------------")
	//创建int类型切片1
	slice1 := []int{1, 2, 3, 4, 5, 6}
	//我们通过整个数组创建一个切片
	slice2 := arr1[:]
	//基于数组的索引下标创建切片
	slice3 := arr2[0:4]
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println("-------------go语言中切片数组当做参数传递-------------------")
	//数组作为参数传递的时候传递的是数组的副本  切片作为函数参数传递的时候传递的实际上是切片的引用
	func1 := func(slice []int) {
		slice[0] = 100
	}
	func2 := func(arr [6]int) {
		arr[0] = 100
	}
	testSlice := []int{1, 2, 3, 4, 5, 6}
	testArray := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(testSlice)
	fmt.Println(testArray)
	func1(testSlice)
	func2(testArray)
	fmt.Println(testSlice)
	fmt.Println(testArray)
	fmt.Println("-------------go语言中切片数组的遍历与迭代-------------------")
	//这里定义了一个数组 我们还可以将它改成切片
	var tmpArr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for index, item := range tmpArr {
		fmt.Println("Index=", index, " Item=", item)
	}
	for index := range tmpArr {
		fmt.Println("Index=", index, " Item=", tmpArr[index])
	}
	for i := 0; i < len(tmpArr); i++ {
		fmt.Println("Index=", i, " Item=", tmpArr[i])
	}
	fmt.Println("-----------------go语言中的数组切片内置函数的使用---------------------")
	arrCap := cap(tmpArr)
	arrLen := len(tmpArr)
	fmt.Println("arrCap:", arrCap, " arrLen:", arrLen)
	//给切片添加元素
	tmpArr = append(tmpArr, 1, 2, 3, 4)
	arrCap = cap(tmpArr)
	arrLen = len(tmpArr)
	fmt.Println("arrCap:", arrCap, " arrLen:", arrLen)
	//给切片添加切片
	addSlice := []int{6, 7, 8, 9}
	tmpArr = append(tmpArr, addSlice...)
	arrCap = cap(tmpArr)
	arrLen = len(tmpArr)
	fmt.Println("arrCap:", arrCap, " arrLen:", arrLen)
	//给切片添加一个数组 是不被允许的
	//	addArray := [...]int{6, 7, 8, 9}
	//	tmpArr = append(tmpArr, addArray...)
	//	arrCap = cap(tmpArr)
	//	arrLen = len(tmpArr)
	//	fmt.Println("arrCap:", arrCap, " arrLen:", arrLen)
	addArray := [...]int{6, 7, 8, 9}
	tmpArr = append(tmpArr, addArray[:]...)
	arrCap = cap(tmpArr)
	arrLen = len(tmpArr)
	fmt.Println("arrCap:", arrCap, " arrLen:", arrLen)
	//copy函数的使用  他可以将源切片拷贝到目标切片
	dst := []int{1, 2, 3, 4, 5, 6}
	source := []int{0, 0, 0}
	copy(dst[1:], source)
	fmt.Println(dst)
}
