package main

import "fmt"

func main() {
	//if else
	var a int = 1
	var b int = 2
	if a > 1 {
		fmt.Println("a>1")
	} else if a < 1 {
		fmt.Println("a<1")
	} else {
		fmt.Println("a=1")
		goto END
	}
	//switch case
	switch b {
	case 0:
		{
			fmt.Println("b=0")
		}
	case 1:
		{
			fmt.Println("b=1")
		}
	case 2:
		{
			fmt.Println("b=2")
		}
		fallthrough
	case 3, 4, 5, 6:
		{
			fmt.Println("b=3,4,5,6")
		}
	default:
		{
			fmt.Println("b=unknow")
		}
	}
	//不带表达式的switch
	switch {
	case b == 0:
		{
			fmt.Println("b=0")
		}
	case b == 1:
		{
			fmt.Println("b=1")
		}
	case b == 2:
		{
			fmt.Println("b=2")
		}
		fallthrough
	case b == 3 || b == 4 || b == 5 || b == 6:
		{
			fmt.Println("b=3,4,5,6")
		}
	default:
		{
			fmt.Println("b=unknow")
		}
	}

END:
	fmt.Println("goto  end")

}
