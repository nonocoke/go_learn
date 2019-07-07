package main

import (
	. "../geometry"
	"fmt"
	"image/color"
)

//type Point struct { X, Y float64 }

// 内嵌结构体,
// 可以直接认为通过嵌入的字段就是ColoredPoint自身的字段
type ColoredPoint struct {
	Point
	Color color.RGBA
}

func (p ColoredPoint) Distance(q Point) float64 {
	return p.Point.Distance(q)
}

func (p *ColoredPoint) ScaleBy(factory float64) {
	p.Point.ScaleBy(factory)
}

func main() {
	var cp ColoredPoint
	cp.X = 1
	fmt.Println(cp.Point.X)  // 1
	cp.Point.Y = 2
	fmt.Println(cp.Y)  // 2  (Don't use if must)

	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{255, 0, 0, 255}
	var p = ColoredPoint{Point{1, 1}, red}
	var q = ColoredPoint{Point{5, 4}, blue}
	fmt.Println(p.Distance(q.Point))  // 5
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point))  // 10


	// 方法值与方法表达式
	a := Point{1, 2}
	b := Point{4, 6}

	distanceFormp := a.Distance  // method value
	fmt.Println(distanceFormp(b))  // 5
	fmt.Printf("distanceFormp's type is %T\n", distanceFormp)  // func(geometry.Point) float64
	var origin Point  // {0, 0}
	fmt.Println(distanceFormp(origin))  // 2.23606797749979, sqrt(5)

	scaleP := a.ScaleBy  // method value
	scaleP(2)
	fmt.Println(a)  // {2, 4}
	scaleP(5)
	fmt.Println(a)  // {10, 20}

	distance := Point.Distance  // method expression
	fmt.Println(distance(a, b))
	fmt.Printf("distance's type is %T\n", distance)  // func(geometry.Point, geometry.Point) float64

	scale := (*Point).ScaleBy  // method expression
	scale(&a, 2)
	fmt.Println(a)
	fmt.Printf("scale's type is %T\n", scale)  // func(*geometry.Point, float64)
}