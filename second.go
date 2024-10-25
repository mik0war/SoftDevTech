package main

import (
	"fmt"
	"strconv"
	"strings"
)

func mapToCelsius(temp int) float32 {
	return (float32)(temp+32) / 1.8
}

func mapToFahrenheit(temp int) float32 {
	return float32(temp)*1.8 + 32
}

func main() {
	string1 := "25 (Celsius)"
	//Получаем значение градуса из входной строки и приводим к типу int
	temp1, _ := strconv.Atoi(strings.Split(string1, " ")[0])

	fmt.Printf("%f (Farhrenheit)\n", mapToFahrenheit(temp1))

	string2 := "32 (Farhrenheit)"
	//Получаем значение градуса из входной строки и приводим к типу int
	temp2, _ := strconv.Atoi(strings.Split(string2, " ")[0])

	fmt.Printf("%f (Celsius)\n", mapToCelsius(temp2))

}
