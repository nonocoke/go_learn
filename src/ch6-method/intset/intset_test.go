package intset

import (
	"fmt"
	"testing"
)

func TestIntSet_String(t *testing.T) {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String())  // {1 9 144}

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String())  // {9 42}

	x.UnionWith(&y)  // {1 9 42 144}
	fmt.Println(x.String())
	fmt.Println(x.Has(9), x.Has(123))
}

func TestIntSet_Len(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.Len())
}

func TestIntSet_Clear(t *testing.T) {
	var x IntSet
	x.Add(9)
	x.Add(144)
	x.Clear()
	fmt.Println(x.String())
}

func TestIntSet_Remove(t *testing.T) {
	var x IntSet
	x.Add(9)
	x.Add(144)
	x.Remove(9)
	fmt.Println(x.String())
}


func TestIntSet_Copy(t *testing.T) {
	var x IntSet
	var y *IntSet
	x.Add(9)
	x.Add(144)
	y = x.Copy()
	fmt.Printf("%T\t%T\t%s\n",x, y, y.String())
}

func TestIntSet_AddAll(t *testing.T) {
	var x IntSet
	x.Add(4)
	x.Add(5)
	sum := x.AddAll(1, 2, 3)
	fmt.Printf("%d", sum)
}
