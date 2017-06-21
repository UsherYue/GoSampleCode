package main

//#include <stdio.h>
/*
void hello() {
        printf("Hello, Cgo!\n");
}
*/
import "C"

func main() {
	C.hello()
}
