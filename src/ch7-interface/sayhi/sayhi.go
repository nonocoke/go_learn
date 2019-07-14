package main

import "fmt"

type Human struct {
	name string
	age int
	phone string
}

type Student struct {
	Human  // 匿名字段
	school string
	loan float32
}

type Employee struct {
	Human
	company string
	money float32
}

// Human 实现SayHi方法
func (h Human) SayHi() {
	fmt.Printf("Hi I am %s you can call me on %s\n", h.name, h.phone)
}

// Human实现Sing方法
func (h Human) Sing(lyrics string) {
	fmt.Println("La lala la ...", lyrics)
}

// Emoloyee重载Human的SayHi方法
func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s, call me on %s\n", e.name, e.company, e.phone)
}

// Interface Men被Human、Student、Employee实现
// 因为这三个类型都实现了这两个方法
type Men interface {
	SayHi()
	Sing(lyrics string)
}

func main() {
	mike := Student{Human{"Mike", 24, "222-222-xxx"}, "MIT", 0.00}
	paul := Student{Human{"Paul", 24, "111-222-xxx"}, "HHHHH", 0.00}
	sam := Employee{Human{"Sam", 33, "444-222-xxx"},"Golang Inc", 1000}
	Tom := Employee{Human{"Tom", 33, "333-222-xxx"},"Things Inc", 4000}

	// 定义Men类型的变量i
	var i Men

	// i 能存储Student
	i = mike
	fmt.Println("This is Mike, an Student: ")
	i.SayHi()
	i.Sing("AA")

	// i 存储Employee
	i = sam
	fmt.Println("This is Sam, an employee: ")
	i.SayHi()
	i.Sing("CC")

	// 定义了 slice Men
	fmt.Println("Let's use a slice of Men and see what happens")
	x := make([]Men, 3)
	// 三个不同类型的元素， 但是它们都实现了interface同一个接口
	x[0], x[1], x[2] = paul, sam, Tom

	for _, value := range x {
		value.SayHi()
	}
}