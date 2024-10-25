package main

import (
	"fmt"
	"strings"
)

func concatStrings(stringsArr []string) string {

	return strings.Join(stringsArr, " ")

}

func main() {
	strings := []string{"Strings", "array", "to", "concat"}
	stringConcat := concatStrings(strings)

	fmt.Println(stringConcat)
}
