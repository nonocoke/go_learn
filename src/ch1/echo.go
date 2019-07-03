package main

import (
	"fmt"
	"os"
)


// 命令行參數使用
func main(){
	//func 1
	//var s, sep string
	//for i := 1; i < len(os.Args); i++ {
	//	s += sep + os.Args[i]
	//	sep = " "
	//}

	//func 2
	s, sep := "", ""
	for i,arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
		fmt.Println(i+1,":"+s)
	}

	//func 3
	//a str
	//fmt.Println(strings.Join(os.Args[1:]," "))  // haha hah - -

	//func 4
	// a slice
	//fmt.Println(os.Args[1:])  // [haha hah - -]
}

