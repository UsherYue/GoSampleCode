package main

import (
	"fmt"
	"reflect"
)

const a1 = 1               //int
const a2 = 1.1             //float64
const a3 = "hello"         //string
const a4 = complex(1, 1.2) //complex128

const (
	b1 = iota       //0
	b2 = iota       //+1=1
	b3              //2
	b4 = "string"   //iota+1
	b5 = iota * 100 //4*100=400
)

const (
	c1 = iota //0
	c2        //1
)

const d1 = iota //0

const (
	One   = iota + 1 //1
	Two              //2
	Three            //3
	Four             //4
)

const (
	e1 = 1 //1
	e2 = 2 //2
	e3     //2
	e4     //2
)

func main() {
	fmt.Println("a1=", a1, " type:", reflect.TypeOf(a1).Name())
	fmt.Println("a2=", a2, " type:", reflect.TypeOf(a2).Name())
	fmt.Println("a3=", a3, " type:", reflect.TypeOf(a3).Name())
	fmt.Println("a4=", a4, " type:", reflect.TypeOf(a4).Name())
	fmt.Println("------------------------------------------")
	fmt.Println("b1=", b1, " type:", reflect.TypeOf(b1).Name())
	fmt.Println("b2=", b2, " type:", reflect.TypeOf(b2).Name())
	fmt.Println("b3=", b3, " type:", reflect.TypeOf(b3).Name())
	fmt.Println("b4=", b4, " type:", reflect.TypeOf(b4).Name())
	fmt.Println("b5=", b5, " type:", reflect.TypeOf(b5).Name())
	fmt.Println("------------------------------------------")
	fmt.Println("c1=", c1, " type:", reflect.TypeOf(c1).Name())
	fmt.Println("c2=", c2, " type:", reflect.TypeOf(c2).Name())
	fmt.Println("------------------------------------------")
	fmt.Println("d1=", d1, " type:", reflect.TypeOf(d1).Name())
	fmt.Println("------------------------------------------")
	fmt.Println("One=", One, " type:", reflect.TypeOf(One).Name())
	fmt.Println("Two=", Two, " type:", reflect.TypeOf(Two).Name())
	fmt.Println("Three=", Three, " type:", reflect.TypeOf(Three).Name())
	fmt.Println("Four=", Four, " type:", reflect.TypeOf(Four).Name())
	fmt.Println("------------------------------------------")
	fmt.Println("e1=", e1, " type:", reflect.TypeOf(e1).Name())
	fmt.Println("e2=", e2, " type:", reflect.TypeOf(e2).Name())
	fmt.Println("e3=", e3, " type:", reflect.TypeOf(e3).Name())
	fmt.Println("e4=", e4, " type:", reflect.TypeOf(e4).Name())

}
