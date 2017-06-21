package main

import (
	"fmt"
	"time"
)

//#include <stdio.h>
//#include <stdlib.h>
/*
int hello() {
        printf("Hello, Cgo!\n");
		return 1;
}
*/
import "C"

func Srand(i int) {
	C.srand(C.uint(i))
}
func Rand() int {
	return int(C.rand())
}

func Hello() int {
	return int(C.hello())
}

func main() {
	fmt.Println(Hello())
	Srand(int(time.Now().Nanosecond()))
	fmt.Println(Rand())
}
