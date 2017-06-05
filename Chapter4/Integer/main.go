package main

import "fmt"

type Integer int

//这里操作的是指针 对应的地址 会影响到值的本身
func (this *Integer) Plus(num Integer) {
	*this += num
}

//这里操作的实际上是Integer的副本并不会影响到值本身
func (this Integer) Minus(num Integer) {
	this -= num
}

//func (this *Integer) Minus(num Integer) {
//	*this -= num
//}

func main() {

	var i Integer = 1
	fmt.Println("i=", i)
	i.Plus(100)
	fmt.Println("i=", i)
	i.Minus(50)
	fmt.Println("i=", i)

}
