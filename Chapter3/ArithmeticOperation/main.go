package main

import "fmt"

func main() {
	//+ - *  / ++ --  %
	var i int = 0
	fmt.Println("i=", i)
	i = i + 1
	fmt.Println("i=", i)
	i++
	fmt.Println("i=", i)
	i--
	fmt.Println("i=", i)
	i = i * 5
	fmt.Println("i=", i)
	i = i / 5
	fmt.Println("i=", i)
	i = i % 100
	fmt.Println("i=", i)
	//+= *= <<=  >>= %= ^= |=
	i += 5 //i=i+5
	fmt.Println("i=", i)
	i *= 9
	fmt.Println("i=", i)
	i %= 53
	fmt.Println("i=", i)
	i <<= 3
	fmt.Println("i=", i)
	i >>= 1
	fmt.Println("i=", i)
	//8  0x08    0000 1000  ^  0000 1001= 0000 0001
	fmt.Println("8^9=", 8^9)
	//0000 0100   ^  0000 1001=0000 1101=8+4+1=13
	i ^= 9
	fmt.Println(i)
	//0000 1101 | 0000 0001= 0000 1101 =8 +4+1 =13
	i |= 1
	fmt.Println(i)
	// &  |  >>  <<  ^
	//0001 & 0010 =0000
	fmt.Println("1&2=", 1&2)
	//0001 | 0010 =0011
	fmt.Println("1|2=", 1|2)
	//0001 ^ 0010 =0011
	fmt.Println("1^2=", 1^2)
	//0001 << 3 = 1000 = 2的三次方=8
	fmt.Println("1<<3=", 1<<3)
	//1000 >> 3  = 0001
	fmt.Println("8>>3=", 8>>3)

}
