package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

// 哈希表是一种巧妙且实用的数据结构。它是一个无序的key/value对的集合，其中所有的key都是不同的，
// 然后通过给定的key可以在常数时间复杂度内检索、更新或刪除对应的value
// map[K]V   e.g., ages := make(map[string]int)


func test_map() {
	ages := map[string]int{
		"alice":	31,
		"charlie":	34,
	}
	// ==>
	//ages := make(map[string]int)
	//ages["alice"] = 31
	//ages["charlie"] = 34

	delete(ages, "alice")  // remove element ages["alcie"]
	if age, ok := ages["bob"]; !ok { fmt.Println("hello")} else {fmt.Println(age)}
}


func main() {
	//test_map()

	// 统计输入中每个Unicode码出现的次数

	counts := make(map[rune]int)	// counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int	// count of lengths of UTF-8 encodings
	invalid := 0					// count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()  //return rune, nbytes, err
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}

		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\t\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Printf("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}

}
