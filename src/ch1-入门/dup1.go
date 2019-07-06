// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.
package main

import (
	"bufio"
	"fmt"
	"os"
)

// 查找重复的行
func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		// Ctrl + D 结束输入
		counts[input.Text()]++
	}

	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
			// NOTE:
			// %d			int变量
			// %x, %o, %b	分别为16进制、8进制、2进制形式的int
			// %f, %g, %e	浮点数: 3.141593 3.141592653589793 3.141593e+00
			// %t			布尔变量: true 或 false
			// %c			rune (Unicode码), go语言特有的Unicode字符类型
			// %s			string
			// %q			带双引号的字符串"abc" 或 带单引号的 rune 'c'
			// %v			会将任意变量以易读的形式打印出来
			// %T			打印变量的类型
			// %%			字符型百分比标志(%符号本身)
		}
	}
}