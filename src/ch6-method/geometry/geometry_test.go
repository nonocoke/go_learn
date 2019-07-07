package geometry

import (
	"fmt"
	"testing"
)

func TestDistance(t *testing.T) {
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(Distance(p, q))  // 5, function call
}

func TestPoint_Distance(t *testing.T) {
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(Distance(p, q))  // 5, method call
}

func TestPath_Distance(t *testing.T) {
	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance())  // 12
	fmt.Println(PathDistance(perim))  // 12
}

func TestPoint_ScaleBy(t *testing.T) {
	// call 1
	r := &Point{1, 2}
	r.ScaleBy(2)
	fmt.Println(*r)  // {2, 4}

	// call 2
	p := Point{1, 2}
	pptr := &p
	pptr.ScaleBy(2)
	fmt.Println(p)  // {2, 4}

	// call 3
	q := Point{1, 2}
	(&q).ScaleBy(2)
	fmt.Println(q)  // {2, 4}
}
