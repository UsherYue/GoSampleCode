package newpkg

import (
	"fmt"
)

func init() {
	fmt.Println("newpkg init....")
}

//首字母小写的函数或者变量 对外不可见
func testNewPkg() {
	fmt.Println("test new pkg")
}

func TestNewPkg() {
	fmt.Println("test new pkg")
}
