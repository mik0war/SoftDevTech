package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getWordsCount(words string) map[string]int {

	wordsMap := map[string]int{}

	for _, word := range strings.Fields(words) {
		_, ok := wordsMap[word]

		if ok {
			wordsMap[word]++
		} else {
			wordsMap[word] = 1
		}

	}

	return wordsMap
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	words := scanner.Text()

	fmt.Println(getWordsCount(words))
}
