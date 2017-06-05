package main

import (
	"fmt"
)

type Integer int

func (this *Integer) Plus(num Integer) {
	*this += num
}

func (this *Integer) Minus(num Integer) {
	*this -= num
}

type Adder interface {
	Plus(num Integer)
	Minus(num Integer)
}

type Pluser interface {
	Plus(num Integer)
}

func main() {
	var i Integer = 1
	var add Adder = &i
	//接口赋值
	var plus Pluser = add
	//+
	plus.Plus(100)
	fmt.Println("After Add:")
	//这里涉及到接口查询的知识我们在后续讲解
	fmt.Println(*(add.(*Integer)))
	//-
	add.Minus(30)
	fmt.Println("After Minus:")
	fmt.Println(*(add.(*Integer)))

}
