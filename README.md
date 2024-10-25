# Технологии разработки программного обеспечения
### Корнеев А.Д. ПИМО-01-24

### Практическая работа №4
## 1. Задачи на линейное программирование (без условных операторов и циклов)

### Задание 1.1.
Напишите программу, которая принимает целое число и вычисляет сумму его цифр.
```go
package main

import (
	"fmt"
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

func main() {
	fmt.Println(addNumberDigits(2_147_483_647))
}
```

**Результат**

![alt-text](https://github.com/mik0war/SoftDevTech/blob/go_pr4/images/4_1.PNG)

### Задание 1.2.
Напишите программу, которая преобразует температуру из градусов Цельсия в Фаренгейты и обратно. 
```go
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
```

**Результат**

![alt-text](https://github.com/mik0war/SoftDevTech/blob/go_pr4/images/4_2.PNG)


### Задание 1.3.
Напишите программу, которая принимает массив чисел и возвращает новый массив, где каждое число удвоено.
```go
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

```


**Результат**

![alt-text](https://github.com/mik0war/SoftDevTech/blob/go_pr4/images/4_3.PNG)


### Задание 1.4.
Напишите программу, которая принимает несколько строк и объединяет их в одну строку через пробел.
```go
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

```

**Результат**

![alt-text](https://github.com/mik0war/SoftDevTech/blob/go_pr4/images/4_4.PNG)


### Задание 1.5.
Напишите программу, которая вычисляет расстояние между двумя точками в 2D пространстве.
```go
package main

import (
	"fmt"
	"math"
)

func getDistance(x1 float64, y1 float64, x2 float64, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}

func main() {
  var x1, x2, y1, y2 float64

	fmt.Println("Введите координаты первой точки: ")
	fmt.Scan(&x1, &y1)

	fmt.Println("Введите координаты второй точки: ")
	fmt.Scan(&x2, &y2)

	fmt.Println(getDistance(x1, y1, x2, y2))
}
```

**Результат**

![alt-text](https://github.com/mik0war/SoftDevTech/blob/go_pr4/images/4_5.PNG)


## 2. Задачи с условным оператором

### Задание 2.1.
Напишите программу, которая проверяет, является ли введенное число четным или нечетным.
```go
package main

import (
	"fmt"
)

func isOdd(number int) bool {
	return number%2 == 0
}

func main() {
	var num int
	fmt.Println("Введите число")
	fmt.Scan(&num)

	if isOdd(num) {
		fmt.Println("Чётное")
	} else {
		fmt.Println("Нечётное")
	}
}
```

**Результат**

![alt-text](https://github.com/mik0war/SoftDevTech/blob/go_pr4/images/4_6.PNG)


### Задание 2.2.
Напишите программу, которая проверяет, является ли введенный год високосным.
```go
package main

import (
	"fmt"
)

func isLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

func main() {
	var year int
	fmt.Println("Введите число")
	fmt.Scan(&year)

	if isLeapYear(year) {
		fmt.Println("Високосный")
	} else {
		fmt.Println("Не високосный")
	}
}
```

**Результат**

![alt-text](https://github.com/mik0war/SoftDevTech/blob/go_pr4/images/4_7.PNG)


### Задание 2.3.
Напишите программу, которая принимает три числа и выводит наибольшее из них.
```go
package main

import (
	"fmt"
)

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

func main() {
	var num1, num2, num3 int
	fmt.Println("Введите числа")
	fmt.Scan(&num1, &num2, &num3)

	maxNum := findMax(num1, num2, num3)
	fmt.Println(maxNum)
}
```

**Результат**

![alt-text](https://github.com/mik0war/SoftDevTech/blob/go_pr4/images/4_8.PNG)


### Задание 2.4.
Напишите программу, которая принимает возраст человека и выводит, к какой возрастной группе он относится (ребенок, подросток, взрослый, пожилой. В комментариях указать возрастные рамки).
```go
package main

import (
	"fmt"
)

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

func main() {
	var age int
	fmt.Println("Введите возраст")
	fmt.Scan(&age)

	fmt.Println(getAgeGroup(age))
}
```

**Результат**

![alt-text](https://github.com/mik0war/SoftDevTech/blob/go_pr4/images/4_9.PNG)


### Задание 2.5.
Напишите программу, которая проверяет, делится ли число одновременно на 3 и 5.
```go
package main

import (
	"fmt"
)

func isDivided5and3(num int) bool {
	return num%3 == 0 && num%5 == 0
}

func main() {
	var num int
	fmt.Println("Введите число")
	fmt.Scan(&num)

	if isDivided5and3(num) {
		fmt.Println("Делится")
	} else {
		fmt.Println("Не делится")
	}
}
```

**Результат**

![alt-text](https://github.com/mik0war/SoftDevTech/blob/go_pr4/images/4_10.PNG)


## 3. Задачи на циклы
### Задание 3.1.
Напишите программу, которая вычисляет факториал числа.
```go
package main

import (
	"fmt"
)

func factorial(number int) int {
	if number == 0 || number == 1 {
		return 1
	}

	return number * factorial(number-1)
}

func main() {
	number := 5
	fmt.Println(factorial(number))
}
```

**Результат**

![alt-text](https://github.com/mik0war/SoftDevTech/blob/go_pr4/images/4_11.PNG)


### Задание 3.2.
Напишите программу, которая выводит первые "n" чисел Фибоначчи.
```go
package main

import (
	"fmt"
	"os"
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
```

**Результат**

![alt-text](https://github.com/mik0war/SoftDevTech/blob/go_pr4/images/4_12.PNG)


### Задание 3.3.
Напишите программу, которая переворачивает массив чисел.
```go
package main

import (
	"fmt"
  "math/rand/v2"
)

func reverce(array []int) []int {
	for i := 0; i < len(array)/2; i++ {
		temp := array[i]
		array[i] = array[len(array)-1-i]
		array[len(array)-1-i] = temp
	}

	return array
}

func main() {
	var arrSize int
	fmt.Println("Введите количество элементов в массиве: ")
	fmt.Scan(&arrSize)

	array := make([]int, arrSize)

	for i := 0; i < arrSize; i++ {
		array[i] = rand.IntN(100)
	}

	fmt.Println(array)
	fmt.Println(reverce(array))
}
```

**Результат**

![alt-text](https://github.com/mik0war/SoftDevTech/blob/go_pr4/images/4_13.PNG)


### Задание 3.4.
Напишите программу, которая выводит все простые числа до заданного числа.
```go
package main

import (
	"fmt"
)

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

func main() {
	var n int
	fmt.Scan(&n)
	for i := 2; i < n; i++ {
		if isPrime(i) {
			fmt.Print(i, " ")
		}
	}
}

```

**Результат**

![alt-text](https://github.com/mik0war/SoftDevTech/blob/go_pr4/images/4_14.PNG)


### Задание 3.5.
Напишите программу, которая вычисляет сумму всех чисел в массиве.
```go
package main

import (
	"fmt"
	"math/rand/v2"
)

func sumArray(array []int) int {
	var sum int
	for _, element := range array {
		sum += element
	}

	return sum
}

func main() {
	var arrSize int
	fmt.Println("Введите количество элементов в массиве: ")
	fmt.Scan(&arrSize)

	array := make([]int, arrSize)

	for i := 0; i < arrSize; i++ {
		array[i] = rand.IntN(100)
	}
	fmt.Println(array)
	fmt.Println(sumArray(array))
}

```


**Результат**

![alt-text](https://github.com/mik0war/SoftDevTech/blob/go_pr4/images/4_15.PNG)

