package main

import "fmt"

type tree struct {
	value	int
	left, right	*tree
}

// treeSort sorts values in place
func treeSort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// add returns the resulting slice
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}

	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

// 结构体可嵌入、匿名成员
func type_struct() {

	type Point struct {
		X, Y int
	}

	type Circle struct {
		Point  // 匿名成员
		Radius int
	}

	type Wheel struct {
		Circle // 匿名成员, 可以通过Wheel.Radius直接访问Radius的值
		Spokes int
	}
	var w Wheel
	w = Wheel{Circle{Point{8, 8}, 5}, 20}
	//w = Wheel{
	//	Circle: Circle{
	//		Center: Point{X: 8, Y: 8},
	//		Radius: 5,
	//	},
	//	Spokes: 20,  // NOTE: trailing comma necessary here (and at Radius)
	//}
	fmt.Printf("%#v\n", w)

	w.X = 42  // equal to w.circle.point.X = 42
	fmt.Printf("%#v\n", w)


}

func main() {
	q := []int{1, 5, 3, 2, 4}
	treeSort(q)
	fmt.Println(q)  // [1 2 3 4 5]

	type_struct()
}