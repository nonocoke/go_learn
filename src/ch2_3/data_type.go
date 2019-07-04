// data type shows some data's type calc.
package main

import "fmt"

/*	整数类型:
	有符号整数: int8、int16、int32和int64
	无符号整数: uint8、uint16、uint32和uint64
	还有两种对应特定CPU平台机器字大小的有符号 int和无符号整数 uint,大小为32或64bit
	Unicode字符 rune类型是和int32等价的类型，通常表示一个Unicode码，可以互换使用
	byte类型是uint8的等价类型，byte一般用于强调数值是一个原始的数据而不是一个小的整数
	uintptr(整数类型),没有具体的bit大小但足以容纳指针
*/
func main() {
	// 类型越界
	var u uint8 = 255
	fmt.Println(u, u+1, u*u)  // 255 0 1

	var i int8 = 127
	fmt.Println(i, i+1, i*i)  // 127 -128 1

	// bit位操作付
	var x uint8 = 1
	var y uint8 = 1<<1 | 1<<2
	fmt.Printf("%08b\n", x)  // 00000001
	fmt.Printf("%08b\n", y)  // 00000110

	fmt.Printf("%08b\n", x&y)  // 00000000
	fmt.Printf("%08b\n", x|y)  // 00000111
	fmt.Printf("%08b\n", x^y)  // 00000111
	fmt.Printf("%08b\n", x&^y)  // 00000001

	// 类型转换
	f := 3.141 // a float64
	g := int(f)
	fmt.Println(f, g)
	f = 1.99
	fmt.Println(int(f))

	// 进制格式打印
	// %d(10)、%o(8)或%x(16)
	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o)  // 438 666 0666
	X := int64(0xdeadbeef)
	fmt.Printf("%d %#[1]x %#[1]X\n", X)  // 3735928559 0xdeadbeef 0XDEADBEEF

	// 字符使用%c参数打印，或者是用%q参数打印带单引号的字符
	ascii := 'a'
	unicode := '国'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)  // 97 a 'a'
	fmt.Printf("%d %[1]c %[1]q\n", unicode)  // 22269 国 '国'
	fmt.Printf("%d %[1]q\n", newline)  // 10 '\n'
}