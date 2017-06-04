package main

import (
	"Package/newpkg"
	"fmt"
)

func init() {
	fmt.Println("main init....")
}

func main() {
	test()
	newpkg.TestNewPkg()
}
