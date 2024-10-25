package main

import (
	"fmt"
)

func double_int_array(arr [4]int) [4]int {
	arr[0] *= 2
	arr[1] *= 2
	arr[2] *= 2
	arr[3] *= 2

	return arr
}

func main() {
	test_array := [4]int{1, 2, 3, 4}
	double_test_array := double_int_array(test_array)
	fmt.Print(double_test_array)
}
