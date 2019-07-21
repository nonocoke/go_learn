package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 文件权限:
	// os.O_RDONLY: 只读
	// os.O_WRONLY: 只写
	// os.O_CREATE: 创建
	// os.O_TRUNC: 截断, 如果文件已存在, 就将该文件的长度截为0
	outputFile, outputError := os.OpenFile("output.dat", os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Printf("An error occured with file opening or creation\n")
		return
	}
	defer outputFile.Close()

	// 写文件方式 1:
	outputWriter := bufio.NewWriter(outputFile)
	outputString := "hello world!\n"

	for i:=0; i<10; i++ {
		outputWriter.WriteString(outputString)
	}
	outputWriter.Flush()

	// 写文件方式 2: fmt.Fprintf(outputFile, "some write data.\n")
	// 写文件方式 3: outputFile.writeString("some write data.\n")
}