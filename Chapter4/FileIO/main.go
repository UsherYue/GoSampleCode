package main

import (
	"FileIO/helper"
	"fmt"
)

func main() {
	fileIO := helper.NewFileIO("./a.txt")
	//创建一个文件
	fileIO.CreateFile()
	//写入文件内容
	fileIO.WriteFile([]byte("hello,world!"))
	//读取文件内容
	bts, _ := fileIO.ReadFile()
	fmt.Println("读出的内容是:", string(bts))
	//删除文件
	fileIO.DeleteFile()
}
