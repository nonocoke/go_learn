package main

import "fmt"

// 可变参数
func sum(vals...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func main() {
	fmt.Println(sum())  // 0
	fmt.Println(sum(3))  // 3
	fmt.Println(sum(1, 2, 3))  //6

	values := []int{1, 2, 3}
	fmt.Println(sum(values...))  // 6
}