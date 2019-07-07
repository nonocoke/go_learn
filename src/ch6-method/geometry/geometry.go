package geometry

import (
	"math"
)

// Note:
// 1. 不管method的receiver是指针类型还是非指针类型, 都是可以通过指针/非指针类型进行调用的, 编译器会隐式做类型转换
// 2. 在声明一个method的receiver该是指针还是非指针类型时, 需要考虑两方面:
//		a. 该对象本身是否特别大, 如果声明为非指针变量时, 调用会产生一次拷贝
//		b. 如果用指针类型作为receiver, 该指向的始终是一块内存地址

type Point struct { X, Y float64 }

// 函数声明
// traditional function
// call geometry.Distance
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}


// 方法声明
// same thing, but as a method of Point type
// 附加参数p, 叫做方法的接收器(receiver), 早起的面向对象语言遗留下的说法: 调用一个方法称为"向一个对象发送消息"
//
// 方法比之函数的一些好处：方法名可以简短。当我们在包外调用的时候这种好处就会被放大，因为可以使用这个短名字，而可以省略掉包的名字
//
// perim := geometry.Path{{1, 1}, {5, 1}, {5, 4}, {1, 1}}
// fmt.Println(geometry.PathDistance(perim)) // "12", standalone function
// fmt.Println(perim.Distance())             // "12", method of geometry.Path
//
// call Point.Distance [声明的Point, Point类下声明的Point.Distance]
// p.Distance(q) 选择器
func (p Point) Distance(q Point) float64 {
	// 计算每个连接点的线段长度
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

type Path []Point

func PathDistance(path Path) float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}


// 基于指针对象的方法
// 当调用一个函数时, 会对其每一个参数值进行拷贝, 如果一个函数需要更新一个变量, 或函数的其中一个参数实在太大,
// 我们希望能够避免这种默认的拷贝, 此时, 需要用到指针
//
// (*Point).ScaleBy
// 接收器只可能是类型(Point)/(*Point), 若类型本身是指针, 是不允许作为接收器的
func (p *Point) ScaleBy(factory float64) {
	p.X *= factory
	p.Y *= factory
}
