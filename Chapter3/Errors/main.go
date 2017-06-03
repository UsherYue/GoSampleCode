package main

import (
	"errors"
	"fmt"
)

func TestError() error {
	return errors.New("this is a error")
}

func main() {

	if err := TestError(); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("no error")
	}

}
