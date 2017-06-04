package main

import "fmt"

type Point struct {
	x float64
	y float64
	//	name string
}

func (this *Point) GetPoint() (float64, float64) {
	return this.x, this.y
}

func (this *Point) SetPoint(x, y float64) {
	this.x = x
	this.y = y
}

//.......

func main() {
	pt1 := Point{x: 1.1, y: 1.3}
	var pt2 = Point{}
	var pt3 = Point{1.5, 1.6}
	fmt.Println("pt1:", pt1)
	fmt.Println("pt2:", pt2)
	fmt.Println("pt3:", pt3)
	//调用SetPoint方法
	pt1.SetPoint(1, 0)
	pt2.SetPoint(1, 0)
	pt3.SetPoint(1, 0)
	x1, y1 := pt1.GetPoint()
	x2, y2 := pt2.GetPoint()
	x3, y3 := pt3.GetPoint()
	fmt.Println("x1=", x1, "  y1=", y1)
	fmt.Println("x2=", x2, "  y2=", y2)
	fmt.Println("x3=", x3, "  y3=", y3)

}
