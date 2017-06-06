package main

import (
	"fmt"
)

func main() {

	fmt.Println("-------------二维数组--------------")
	//定义一个二维数组
	var twoDimArr1 [5][3]int
	fmt.Println(twoDimArr1)
	//定义一个一维数组 数组每一个成员是[]int空切片
	var twoDimArr2 [5][]int
	twoDimArr2[1] = []int{1, 2, 3}
	fmt.Println(twoDimArr2)
	fmt.Println("-------------三维数组--------------")
	//定义一个三维数组
	var threeDimArr2 [2][3][5]int
	fmt.Println(threeDimArr2)
	//定义一个二维数组,第二维每一项是[]int空切片
	var threeDimArr1 [5][10][]int
	threeDimArr1[0][0] = []int{12, 3, 4, 5}
	fmt.Println(threeDimArr1)
	fmt.Println("-------------数组切片结合--------------")
	//一维切片 数组每一项都是一个[3]int类型数组
	var sliceAndArr [][3]int = [][3]int{[3]int{1, 3, 4}, [3]int{4, 5, 4}}
	fmt.Println(sliceAndArr)
	//二维切片
	var slice [][]int
	fmt.Println(slice)
	slice = [][]int{[]int{1, 2, 3}, []int{6, 6, 6}}
	fmt.Println(slice)
	fmt.Println("-------------三维数组遍历--------------")
	//三维数组遍历 threeDimArr3  定义并且初始化
	var threeDimArr3 [3][3][2]int = [3][3][2]int{[3][2]int{[2]int{1, 2}}}
	for i := 0; i < len(threeDimArr3); i++ {
		for j := 0; j < len(threeDimArr3[i]); j++ {
			for k := 0; k < len(threeDimArr3[i][j]); k++ {
				fmt.Println("threeDimArr3[", i, "][", j, "][", k, "]=", threeDimArr3[i][j][k])
			}
		}
	}
}
