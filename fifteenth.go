package main

import (
	"fmt"
	"sort"
)

func createRomanMapAndKeyList() (map[int]string, []int) {
	romanMap := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}

	keys := make([]int, 0, len(romanMap))
	for k := range romanMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for i := 0; i < len(keys)/2; i++ {
		temp := keys[i]
		keys[i] = keys[len(keys)-1-i]
		keys[len(keys)-1-i] = temp
	}

	return romanMap, keys
}

func toRomanNumbers(number int) string {

	romanMap, keys := createRomanMapAndKeyList()

	result := ""

	for number > 0 {
		for _, key := range keys {
			for number/key != 0 {
				result += romanMap[key]
				number -= key
			}
		}
	}

	return result
}

func main() {

	var num int
	fmt.Println("Введите число:")
	fmt.Scan(&num)

	fmt.Print(toRomanNumbers(num))
}
