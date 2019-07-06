package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

/*
	ASCII: 	 7bit 来表示128个字符(包含英文字母的大小写、数字、各种标点符号和设置控制符)
	Unicode: 收集了所有符号系统, 包括重音符号和其他变音符号、制表符和回车符等, 每个符号都被分配一个唯一的Unicode码,Unicode码对应Go中
			 的rune整数类型(int32)
	UTF-8:	 一种将Unicode码编码为字节序列的变长编码。
			 UTF-8编码使用1到4个字节来表示每个Unicode码, ASCII部分字节只使用1个字节, 常用字符部分使用2或3个字节表示。
	------------------
	标准库 - 字符串处理包
	strings: 提供了许多如字符串的查询、替换、比较、截断、拆分和合并等功能。
	bytes:	 提供了很多类似功能的函數，但是针对的是和字符串有着相同结构的[]byte类型   (byte.Buffer类型)。
	strconv: 提供了布尔型、整型数、浮点数和对应字符串的相互转换，还提供了双引号转义相关的转换。
	unicode: 提供了IsDigit、IsLetter、IsUpper和IsLower等类似功能，用于给字符分类。
			 每個函数有一个单一的rune类型的参数，然后返回一個布尔值
 */

// Strings.HasPrefix
func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

// Strings.HasSuffix
func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s) - len(suffix):] == suffix
}

// Strings.Contains
func Contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}

func testCode() {
	temp := "世界"
	temp1 := "\xe4\xb8\x96\xe7\x95\x8c"
	temp2 := "\u4e16\u754c"  // \uhhhh 对应16bit的字节码
	temp3 := "\U00004e16\U0000754c"  // \Uhhhhhhhh对应32bit的字节码

	fmt.Println(temp, temp1, temp2, temp3)

	//fmt.Println(strings.HasPrefix(temp, temp1))  // 判断temp1 是否为 temp 的前缀
	//fmt.Println(strings.HasSuffix(temp, temp1))  // 判断temp1 是否为 temp 的后缀
	//fmt.Println(strings.Contains(temp, temp1))  // 判断temp1 是否为 temp 的子串
	fmt.Println(HasPrefix(temp, temp1))
	fmt.Println(HasSuffix(temp, temp1))
	fmt.Println(Contains(temp, temp1))

	s := "hello, 世界"   // 字符串s 包含有13个字节, 以UTF8形式编码, 只对应9个Unicode字符
	fmt.Println(len(s))  // 13
	fmt.Println(utf8.RuneCountInString(s))  // 9

	// UTF8解码器  unicode/utf8.DecodeRuneInString()
	//for i := 0; i < len(s); {
	//	r, size := utf8.DecodeRuneInString(s[i:])
	//	fmt.Printf("%d\t%c\n", i, r)
	//	i += size
	//}

	// Go语言的range循环在处理字符串的时候，会自动隐式解码UTF8字符串
	for i, r := range "hello, 世界" {
		fmt.Printf("%d\t%c\n", i, r)
	}

	// 统计字符串中字符的数目
	//n := 0
	////for _, _  = range s {
	//for range s {
	//	n++
	//}
	//fmt.Println(n)

	// UTF8字符串作为交换格式是非常方便的，但是在程序內部采用rune序列可能更方便，因为rune大小一致，支持数组索引和方便切割
	ss := "プログラム"  // program <japanese>
	fmt.Printf("% x\n", ss)  // "% x"参数用于在每個十六进制数字前插入一個空格
	rr := []rune(ss)
	fmt.Printf("%x\n", rr)  // [30d7 30ed 30b0 30e9 30e0]
	fmt.Println(string(rr))  // プログラム

	fmt.Println(string(65))  // A
	fmt.Println(string(1234567))  // (?)  对应字节码的字符是无效的，则用'\uFFFD'无效字符作为替换
}

// basename(s)将看起來像是系统路径的前缀刪除，同時将看似文件类型的后缀名部分刪除
// e.g., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
func basename0(s string) string {
	//Discard last '/' and everything before.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1 :]
			break
		}
	}

	// Preserve everything before last '.'.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

// 使用strings.LastIndex()
func basename1(s string) string {
	slash := strings.LastIndex(s, "/")  // -1 if "/" not found
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

// path和path/filepath包提供了关于文件路径名相关函数操作
func testBasename() {
	var s string
	s = basename0("hello/1.go")
	fmt.Println(s)
	s = basename0("/1.hello.go")
	fmt.Println(s)
	s = basename1("/1.hello.go")
	fmt.Println(s)
}

func intChangestyle(s string) string {
	// 判断s 是否为整型
	_, err := strconv.Atoi(s)
	if err != nil {
		return s
	}
	n := len(s)
	if n <= 3 {
		return s
	}
	return intChangestyle(s[:n-3]) + "," + s[n-3:]
}

func testIntChangestyle() {
	var s string
	s = intChangestyle("12345.678")
	fmt.Println(s)
	s = intChangestyle("12345678")
	fmt.Println(s)
}

func main() {

	testCode()
	testBasename()
	testIntChangestyle()
}
