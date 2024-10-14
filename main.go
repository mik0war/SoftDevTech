package main

import (
	"fmt"
	"math"
	"math/rand/v2"
	"strconv"
	"strings"
)

func addNumberDigits(number int32) int32 {
	var sum int32
	//Максимальное количество цифр в числе типа int32 - 10,
	//поэтому повторяем операцию 10 раз, чтобы не использовать цикл
	sum += number % 10
	number /= 10

	sum += number % 10
	number /= 10

	sum += number % 10
	number /= 10

	sum += number % 10
	number /= 10

	sum += number % 10
	number /= 10

	sum += number % 10
	number /= 10

	sum += number % 10
	number /= 10

	sum += number % 10
	number /= 10

	sum += number % 10
	number /= 10

	sum += number % 10
	number /= 10

	return sum
}

func mapToCelsius(temp int) float32 {
	return (float32)(temp+32) / 1.8
}

func mapToFahrenheit(temp int) float32 {
	return float32(temp)*1.8 + 32
}

func doubleIntArray(arr [4]int) [4]int {
	arr[0] *= 2
	arr[1] *= 2
	arr[2] *= 2
	arr[3] *= 2

	return arr
}

func concatStrings(stringsArr []string) string {

	return strings.Join(stringsArr, " ")

}

func getDistance(x1 float64, y1 float64, x2 float64, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}

func isOdd(number int) bool {
	return number%2 == 0
}

func isLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

func findMax(num1 int, num2 int, num3 int) int {
	isFirstBiggerSecond := num1 > num2
	isFirstBiggerThird := num1 > num3
	isSecondBiggerThird := num2 > num3

	if isFirstBiggerSecond && isFirstBiggerThird {
		return num1
	}

	if !isFirstBiggerSecond && isSecondBiggerThird {
		return num2
	}

	return num3
}

func getAgeGroup(age int) string {
	//ребенок 0 - 13, подросток 14 - 21, взрослый 22 - 59, пожилой >60
	if age >= 0 && age < 13 {
		return "ребенок"
	}

	if age >= 14 && age < 21 {
		return "подросток"
	}

	if age >= 22 && age < 59 {
		return "взрослый"
	}

	if age >= 60 {
		return "пожилой"
	}

	return "Некорректное значение"
}

func isDivided5and3(num int) bool {
	return num%3 == 0 && num%5 == 0
}

func factorial(number int) int {
	if number == 0 || number == 1 {
		return 1
	}

	return number * factorial(number-1)
}

func fibonachiNumber(num int) int {
	if num == 0 || num == 1 {
		return num
	}

	return fibonachiNumber(num-1) + fibonachiNumber(num-2)
}

func sumArray(array []int) int {
	var sum int
	for _, element := range array {
		sum += element
	}

	return sum
}

func isPrime(num int) bool {
	if num == 2 {
		return true
	}
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func reverce(array []int) []int {
	for i := 0; i < len(array)/2; i++ {
		temp := array[i]
		array[i] = array[len(array)-1-i]
		array[len(array)-1-i] = temp
	}

	return array
}

func main() {

	//1.1
	fmt.Println(addNumberDigits(2_147_483_647))

	//1.2
	string1 := "25 (Celsius)"
	temp1, _ := strconv.Atoi(strings.Split(string1, " ")[0])

	fmt.Printf("%f (Farhrenheit)\n", mapToFahrenheit(temp1))

	string2 := "32 (Farhrenheit)"
	temp2, _ := strconv.Atoi(strings.Split(string2, " ")[0])

	fmt.Printf("%f (Celsius)\n", mapToCelsius(temp2))

	//1.3

	test_array := [4]int{1, 2, 3, 4}
	double_test_array := doubleIntArray(test_array)
	fmt.Println(double_test_array)

	//1.4
	strings := []string{"Strings", "array", "to", "concat"}
	stringConcat := concatStrings(strings)

	fmt.Println(stringConcat)

	//1.5
	var x1, x2, y1, y2 float64

	fmt.Println("Введите координаты первой точки: ")
	fmt.Scan(&x1, &y1)

	fmt.Println("Введите координаты второй точки: ")
	fmt.Scan(&x2, &y2)

	fmt.Println(getDistance(x1, y1, x2, y2))

	//2.1
	var num int
	fmt.Println("Введите число")
	fmt.Scan(&num)

	if isOdd(num) {
		fmt.Println("Чётное")
	} else {
		fmt.Println("Нечётное")
	}

	//2.2
	var year int
	fmt.Println("Введите число")
	fmt.Scan(&year)

	if isLeapYear(year) {
		fmt.Println("Високосный")
	} else {
		fmt.Println("Не високосный")
	}

	//2.3
	var num1, num2, num3 int
	fmt.Println("Введите числа")
	fmt.Scan(&num1, &num2, &num3)

	maxNum := findMax(num1, num2, num3)
	fmt.Println(maxNum)

	//2.4
	var age int
	fmt.Println("Введите возраст")
	fmt.Scan(&age)

	fmt.Println(getAgeGroup(age))

	//2.5
	var num53 int
	fmt.Println("Введите число")
	fmt.Scan(&num53)

	if isDivided5and3(num53) {
		fmt.Println("Делится")
	} else {
		fmt.Println("Не делится")
	}

	//3.1
	number := 5
	fmt.Println(factorial(number))

	//3.2
	var n int
	fmt.Scan(&n)

	for i := 0; i <= n; i++ {
		fmt.Print(fibonachiNumber(i), " ")
	}

	//3.3
	var arrSize int
	fmt.Println("Введите количество элементов в массиве: ")
	fmt.Scan(&arrSize)

	array := make([]int, arrSize)

	for i := 0; i < arrSize; i++ {
		array[i] = rand.IntN(100)
	}

	fmt.Println(array)
	fmt.Println(reverce(array))

	//3.4
	fmt.Scan(&n)
	for i := 2; i < n; i++ {
		if isPrime(i) {
			fmt.Print(i, " ")
		}
	}

	//3.5
	fmt.Println("Введите количество элементов в массиве: ")
	fmt.Scan(&arrSize)

	array = make([]int, arrSize)

	for i := 0; i < arrSize; i++ {
		array[i] = rand.IntN(100)
	}
	fmt.Println(array)
	fmt.Println(sumArray(array))
}
