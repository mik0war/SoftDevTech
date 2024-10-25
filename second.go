package main

import (
	"fmt"
)

func main() {
	var a uint
	fmt.Println("Введите число a")
	fmt.Scan(&a)

	var b uint
	fmt.Println("Введите число b")
	fmt.Scan(&b)

	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}

	fmt.Print(a)
}
