package main

import "fmt"

type Point struct {
	x float64
	y float64
}

//带指针
func (this *Point) SetPoint1(x, y float64) {
	this.x = x
	this.y = y
}

//不带指针
func (this Point) SetPoint2(x, y float64) {
	this.x = x
	this.y = y
}

func main() {

	//定义一个结构体指针
	pt := new(Point)
	fmt.Println(pt)
	pt.SetPoint1(100, 100)
	fmt.Println(pt)
	pt.SetPoint2(0, 0)
	fmt.Println(pt)

}
