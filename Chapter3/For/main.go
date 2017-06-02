package main

import "fmt"

func main() {

	//利用for实现其他语言中的while(true){}的效果
	a := 1 //var a int =1
	for {
		if a < 100 {
			//fmt.Println(a)
			a++
		} else {
			break
		}
	}
	//传统for循环 也就是带条件的for循环
	for i := 0; i < 100; i++ {
		fmt.Println("i=", i)
	}
	//省略条件和初始值的for循环
	j := 0 // var j int =0
	for ; ; /*j < 100*/ j++ {
		if j < 100 {
			fmt.Println("j=", j)
		} else {
			break
		}
	}
	//for range 用法 去遍历 map字典  或者数组
	var arr []int = []int{1, 2, 3, 4, 5, 6, 7}
	for index, item := range arr {
		fmt.Println("Index=", index, " Item=", item)
	}

	for index := range arr {
		fmt.Println("Index=", index, " Item=", arr[index])
	}

	//map的循环
	mp := map[string]string{"zhangsan": "1", "lisi": "2"}
	for k, v := range mp {
		fmt.Println("K=", k, " V=", v)
	}

	for k := range mp {
		fmt.Println("K=", k, " V=", mp[k])
	}

}
