package main

import "fmt"

func main() {
	//定义一个包含中英文的字符串
	str := "hello,go语言就业指南"
	//我们开始遍历一下字符串
	fmt.Println("通过字符串下标索引每一个字符:")
	//也就是说这种遍历方式并没有遍历出字符 而是将字符串的每个字节给遍历出来了
	for i := 0; i < len(str); i++ {
		fmt.Printf("%d ", str[i])
	}
	fmt.Println("")
	//我们将字符串转换成字节数组 然后遍历字节数组，和我们直接通过len获取字符串的长度然后通过索引下标去遍历字符串是一样的 本质上操作的是字节数组 而不是字符数组
	btArr := []byte(str)
	for i := 0; i < len(btArr); i++ {
		fmt.Printf("%d ", btArr[i])
	}
	fmt.Println("")
	//for range遍历字符    index是每一个字符对应的字节数组的起始下标 code 对应的是每一个utf8 字符类型的编码值
	for index, code := range str {
		//fmt.Println("index:", index, " char:", code)
		fmt.Printf("index=%d,char=%#q\n", index, code)
	}
	//将字符串转换成rune类型的数组切片然后遍历
	runeArr := []rune(str)
	for index, code := range runeArr {
		fmt.Printf("index=%d,char=%#q\n", index, code)
	}

	//获取字符串的某一个字符
	fmt.Printf("最后一个字符是:%#q\n", runeArr[len(runeArr)-1])
	fmt.Printf("第一个字符是:%#q\n", runeArr[0])

	//获取字符串中的子串
	fmt.Println("前五个字符是:", string(runeArr[0:5]))
	fmt.Println("第三个字符往后所有的字符:", string(runeArr[3:]))

}
