package main

import (
	"crypto/sha256"
	"fmt"
)

func testArray() {

	var a [3]int  // default 0
	fmt.Println(a[0])  // print the first element
	fmt.Println(a[len(a)-1])  // print the last element

	// print the indices and elements
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	// print the elements only
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	// initilize
	//var q [3]int = [3]int{1, 2, 3}3
	q := [...]int{1, 2, 3}  // 简化定义
	//var r [3]int = [3]int{1, 2}
	r := [...]int{3: -1}  // 定义含有4个元素的数组r，最后一个元素被初始化为-1
	fmt.Printf("%T\n", q)  // [3]int
	fmt.Println(q) // 1, 2, 3
	fmt.Println(r) // 1, 2, 0

}

func main() {
	//testArray()

	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	// Output:
	// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
	// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	// false
	// [32]uint8
}