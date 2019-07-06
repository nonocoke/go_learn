package main

import (
	"fmt"
)

// Slice（切片）代表变长的序列，序列中每个元素都有相同的类型。
// 一個slice类型一般协作[]T，其中T代表slice中元素的类型；slice的语法和数组很像，知识沒有固定长度而已
// 一个slice由三部分构成: 指针、长度和容量
// 		make([]T, len)
// 		make([]T, len, cap)  // same as make([]T, cap)[:len]


/* slice 模拟实现栈stack

stack = append(stack, v)  // push v
top := stack[len(stack)-1]  // top of stack
stack = stack[:len(stack)-1]  // pop

// 删除slice中某个元素并保持原有的元素顺序, 通过内置的copy函数将后面的子slice向前依次移动一位完成
func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

// 如果删除元素后不用保持原来顺序, 可以简单的用最后一个元素覆盖被删除的元素
func remove(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}
*/

func test_slice() {
	months := [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June",
		7: "July", 8: "August", 9: "September", 10: "October", 11: "Novembr", 12: "December"}
	fmt.Printf("%s\n", months)

	Q2 := months[4:7]
	fmt.Printf("The second quarter is %s\n", Q2)
	summer := months[6:9]
	fmt.Printf("The summer include %s\n", summer)
	fmt.Printf("The summer'len is  %d\n", len(summer))
	fmt.Printf("The summer'cap is  %d\n", cap(summer))
	//fmt.Println(summer[:20])  // anic: runtime error: slice bounds out of range
	endlessSummer := summer[:len(months)-6]
	fmt.Printf("%s\n", endlessSummer)

}

// 因为slice值包含指向第一个slice元素的指针，因此向函數传递slice将允许在函数内部修改底层数组的元素。
// 换句话说，复制一个slice只是对底层的数组创建了一个新的slice别名
// slice反转
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func test_reverse() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	//fmt.Println(a)

	// 将a向左旋转n个元素的方法: 调用3次reverse反转函数
	// 反转开头的n个元素， 然后反转剩下的元素，最后反转整个slice的元素
	s := []int{0, 1, 2, 3, 4, 5}
	reverse(s[:2])
	fmt.Println(s)
	reverse(s[2:])
	fmt.Println(s)
	reverse(s)
	fmt.Println(s)
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// There is room to grow. Extend the slice.
		z = x[:zlen]
	} else {
		// There is insufficient spcae. Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}

		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

func test_appendInt() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(x), x)
		x = y
	}
}

// 去除slice中 值为"" 的元素
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func test_nonempty() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data))
	// 输入的slice和输出的slice共享一个底层数组。这可以避免分配另一个数组，不过原來的数据将可能会被覆盖
	// 用法: data = nonempty(data)
	fmt.Printf("%q\n", data)
}


func main() {
	//test_slice()
	//test_reverse()
	//test_appendInt()
	test_nonempty()

	// 内置的append方法，添加元素或slice
	var x []int
	x = append(x, 1)
	x = append(x, 2, 3)
	x = append(x, 4, 5, 6)
	x = append(x, x...)
	fmt.Println(x)

}