package main

import (
	"fmt"
	"reflect"
)

func main() {
	//complex64
	var cmplx1 complex64 = complex(1.1, 2.2)
	fmt.Println("cmplx1=", cmplx1, " type=", reflect.TypeOf(cmplx1).Name())
	//complex128
	var cmplx2 complex128 = complex(1.1, 2.3)
	fmt.Println("cmplx2=", cmplx2, " type=", reflect.TypeOf(cmplx2).Name())
	fmt.Println("cmplx1*3=", cmplx1*3)
	fmt.Println("cmplx1*cmplx2=", complex128(cmplx1)*cmplx2)
}
