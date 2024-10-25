package main

import (
	"fmt"
)

func fibonachiNumber(num int) int {
	if num == 0 || num == 1 {
		return num
	}

	return fibonachiNumber(num-1) + fibonachiNumber(num-2)
}

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i <= n; i++ {
		fmt.Print(fibonachiNumber(i), " ")
	}
}
