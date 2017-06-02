package main

import "fmt"

//带参数不带返回值的函数
func PrintInfo(name string, age int, addr string) {
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Addr:", addr)
}

//带参数  并且带返回值的函数
func RetInfo(name string, age int, addr string) (string, int, string) {
	return name, age, addr
}

//不定参数
func DelUsers(uids ...int) {
	fmt.Println("Uids:", uids)
}

func main() {
	PrintInfo("张三", 15, "山东")
	name, _, addr := RetInfo("张三", 16, "北京")
	fmt.Println("Name:", name)
	//fmt.Println("Age:", age)
	fmt.Println("Addr:", addr)
	//调用带不定参数的函数
	DelUsers(1, 2, 3, 4, 5, 6, 7)

}
