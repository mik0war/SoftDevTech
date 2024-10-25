# Технологии разработки программного обеспечения
### Корнеев А.Д. ПИМО-01-24

### Практическая работа №5

# Задачи по языку программирования Go
1. **Проверка на простоту**
   
Напишите функцию, которая проверяет, является ли переданное число простым. Ваша программа должна использовать циклы для проверки делителей,
и если число не является простым, выводить первый найденный делитель.
```go
package main

import (
	"fmt"
)

func isPrime(number uint) uint {
	if number <= 1 {
		return 1
	}
	for i := uint(2); i < number; i++ {
		if number%i == 0 {
			return i
		}
	}

	return 0
}

func main() {
	var num uint
	fmt.Println("Введите число")
	fmt.Scan(&num)

	isPrime := isPrime(num)

	switch isPrime {
	case 0:
		fmt.Println("Число простое")
		break
	case 1:
		fmt.Println("Не простое (1 или 0)")
		break

	default:
		fmt.Printf("Число составное, наименьший делитель - %d", isPrime)
	}
}

```

2. **Наибольший общий делитель (НОД)**

Напишите программу для нахождения наибольшего общего делителя (НОД) двух чисел с использованием алгоритма Евклида. Используйте цикл `for` для вычислений

```go
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

```

3. **Сортировка пузырьком**
   
Реализуйте сортировку пузырьком для списка целых чисел. Программа должна выполнять сортировку на месте и выводить каждый шаг изменения массива

```go
package main

import (
	"fmt"
	"math/rand/v2"
)

func bubbleSort(array []int) []int {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-1; j++ {
			if array[j] < array[j+1] {
				temp := array[j]
				array[j] = array[j+1]
				array[j+1] = temp
			}
			fmt.Println(array)
		}
	}

	return array
}

func main() {

	var n int
	fmt.Println("Введите размер массива")
	fmt.Scan(&n)

	array := make([]int, n)

	for i := 0; i < len(array); i++ {
		array[i] = rand.IntN(100)
	}

	fmt.Println("Source array: ", array)
	fmt.Println("Sorted array: ", bubbleSort(array))

}

```

4. **Таблица умножения в формате матрицы** 

Напишите программу, которая выводит таблицу умножения в формате матрицы 10x10. Используйте циклы для генерации строк и столбцов

```go
package main

import (
	"fmt"
	"strings"
)

func mulMatrix(size int) [][]int {
	matrix := make([][]int, size)

	for i := range matrix {
		matrix[i] = make([]int, size)
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			matrix[i][j] = (i + 1) * (j + 1)
		}
	}

	return matrix
}

func main() {

	var n int
	fmt.Println("Введите размер таблицы умножения")
	fmt.Scan(&n)
	matrix := mulMatrix(n)
	for i := 0; i < n; i++ {
		fmt.Println(strings.Trim(strings.Replace(fmt.Sprint(matrix[i]), " ", "\t", -1), "[]"))
	}

}

```

5. **Фибоначчи с мемоизацией**

Напишите функцию для вычисления числа Фибоначчи с использованием мемоизации (сохранение ранее вычисленных результатов). Программа должна использовать рекурсию и условные операторы.

```go
package main

import (
	"fmt"
)

var values = map[int]int{}

func fibonacci(n int) int {
	if n == 1 || n == 0 {
		return 1
	}

	val, ok := values[n]
	if ok {
		return val
	}

	values[n] = fibonacci(n-1) + fibonacci(n-2)
	return values[n]
}

func main() {

	var n int
	for true {
		fmt.Println("Введите номер числа фибоначчи:")
		fmt.Scan(&n)

		fmt.Println("Число фиббоначи: ", fibonacci(n))
		fmt.Println("Сохранённые значения: ", values)
	}

}

```

6. **Обратные числа**

Напишите программу, которая принимает целое число и выводит его в обратном порядке. Например, для числа 12345 программа должна вывести 54321. Используйте цикл для обработки цифр числа.

```go
package main

import (
	"errors"
	"fmt"
)

func reverceNumber(n int) int {
	if n < 0 {
		errors.New("Illegal argument")
	}

	result := 0

	for n > 0 {
		result = result*10 + (n % 10)
		n /= 10
	}

	return result
}

func main() {

	var n int
	for true {
		fmt.Println("Введите число:")
		fmt.Scan(&n)

		fmt.Println("Перевёрнутое число: ", reverceNumber(n))
	}

}

```
7. **Треугольник Паскаля**

Напишите программу, которая выводит треугольник Паскаля до заданного уровня. Для этого используйте цикл и массивы для хранения предыдущих значений строки треугольника.

```go
package main

import (
	"fmt"
)

func pascalTriangle(level int) {
	previousLine := []int{}
	currentLine := []int{}

	for i := 0; i < level; i++ {
		currentLine = append(currentLine, 1)

		for i := 1; i < len(currentLine)-1; i++ {
			currentLine[i] = previousLine[i-1] + previousLine[i]
		}

		fmt.Println(currentLine)
		previousLine = []int{}
		previousLine = append(previousLine, currentLine...)
	}

}

func main() {

	for true {
		var n int
		fmt.Println("Введите уровень треугольника Паскаля")
		fmt.Scan(&n)

		pascalTriangle(n)
	}

}

```

8. **Число палиндром**

Напишите программу, которая проверяет, является ли число палиндромом (одинаково читается слева направо и справа налево). Не используйте строки для решения этой задачи — работайте только с числами.

```go
package main

import (
	"errors"
	"fmt"
)

func reverceNumber(n int) int {
	if n < 0 {
		errors.New("Illegal argument")
	}

	result := 0

	for n > 0 {
		result = result*10 + (n % 10)
		n /= 10
	}

	return result
}

func isPalindrome(number int) bool {
	return number == reverceNumber(number)
}

func main() {

	for true {
		var n int
		fmt.Println("Введите число")
		fmt.Scan(&n)

		if isPalindrome(n) {
			fmt.Println("Палиндром")
		} else {
			fmt.Println("Не палиндром")
		}
	}

}

```

9. **Нахождение максимума и минимума в массиве**

Напишите функцию, которая принимает массив целых чисел и возвращает одновременно максимальный и минимальный элемент с использованием одного прохода по массиву.

```go
package main

import (
	"fmt"
	"math/rand/v2"
)

func minMax(array []int) (int, int) {

	//Помещаем максимальное значение int в переменную min
	min := int(^uint(0) >> 1)
	//Помещаем минимальное значение int в переменную max
	max := -min - 1

	for _, element := range array {

		if element > max {
			max = element
		}

		if element < min {
			min = element
		}
	}

	return min, max
}

func main() {

	for true {
		var n int
		fmt.Println("Введите размер массива")
		fmt.Scan(&n)

		array := make([]int, n)

		for i := 0; i < n; i++ {
			array[i] = rand.IntN(1000) - 500
		}
		fmt.Println("Source array: ", array)
		fmt.Println(minMax(array))
	}

}

```

10. **Игра "Угадай число"**

Напишите программу, которая загадывает случайное число от 1 до 100, а пользователь пытается его угадать. Программа должна давать подсказки \"больше\" или \"меньше\" после каждой попытки. Реализуйте ограничение на количество попыток.

```go
package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {

	var attempts int
	fmt.Println("Количество попыток:")
	fmt.Scan(&attempts)

	n := rand.IntN(101) + 1
	var choice int

	for i := 0; i < attempts; i++ {
		fmt.Println("Попытка: ")
		fmt.Scan(&choice)

		if choice == n {
			fmt.Println("Ты выиграл, число - ", n)
			return
		}

		if choice > n {
			fmt.Println("Загаданное число меньше чем ", choice)
			continue
		}

		fmt.Println("Загаданное число больше чем ", choice)
	}

	fmt.Println("Ты проиграл( Загадано было число ", n)

}

```

11. **Числа Армстронга**

Напишите программу, которая проверяет, является ли число числом Армстронга (число равно сумме своих цифр, возведённых в степень, равную количеству цифр числа). Например, 153 = 1³ + 5³ + 3³.

```go
package main

import (
	"fmt"
	"math"
)

func getDigits(num int) []int {
	result := []int{}

	for num > 0 {
		result = append(result, num%10)
		num /= 10
	}

	return result
}

func checkIsArmstrong(num int, digits []int) bool {
	pow := len(digits)
	digitsSum := 0
	for _, element := range digits {
		digitsSum += int(math.Pow(float64(element), float64(pow)))
	}

	return num == digitsSum
}

func main() {

	num := 0
	fmt.Scan(&num)

	if checkIsArmstrong(num, getDigits(num)) {
		fmt.Println("Число Армстронга")
	} else {
		fmt.Print("Не число Армстронга")
	}
}

```

12. **Подсчет слов в строке**

Напишите программу, которая принимает строку и выводит количество уникальных слов в ней. Используйте `map` для хранения слов и их количества.

```go
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

```

13. **Игра "Жизнь" (Conway's Game of Life)** 

Реализуйте клеточный автомат "Жизнь" Конвея для двухмерного массива. Каждая клетка может быть либо живой, либо мертвой. На каждом шаге состояния клеток изменяются по следующим правилам:
- Живая клетка с двумя или тремя живыми соседями остаётся живой, иначе умирает.
- Мёртвая клетка с тремя живыми соседями оживает.
Используйте циклы для обработки клеток.

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func createMatrix(matrixSize int) [][]int {

	matrix := make([][]int, matrixSize)

	for i := range matrix {
		matrix[i] = make([]int, matrixSize)
	}

	return matrix
}

func getAliveNeighbours(x int, y int, matrix [][]int) int {
	result := 0
	for neighbourX := x - 1; neighbourX <= x+1; neighbourX++ {
		for neighbourY := y - 1; neighbourY <= y+1; neighbourY++ {
			if neighbourX < 0 || neighbourY < 0 || neighbourX >= len(matrix) ||
				neighbourY >= len(matrix[0]) || (neighbourX == x && neighbourY == y) {
				continue
			}
			if matrix[neighbourX][neighbourY] == 1 {
				result++
			}
		}
	}

	return result
}

func showMatrix(matrix [][]int) {
	for _, line := range matrix {
		fmt.Println(line)
	}
}

func makeStep(matrixSize int, matrix [][]int) [][]int {
	result := createMatrix(matrixSize)
	for x, line := range matrix {
		for y, cell := range line {
			result[x][y] = matrix[x][y]
			aliveNeighbours := getAliveNeighbours(x, y, matrix)
			if cell == 1 && !(aliveNeighbours == 3 || aliveNeighbours == 2) {
				result[x][y] = 0
				continue
			}

			if cell == 0 && aliveNeighbours == 3 {
				result[x][y] = 1
				continue
			}
		}
	}

	return result
}
func main() {

	const fieldSize = 10

	matrix := createMatrix(fieldSize)

	matrix[1][4] = 1
	matrix[2][2] = 1
	matrix[2][4] = 1
	matrix[3][3] = 1
	matrix[3][4] = 1

	for true {

		showMatrix(matrix)

		fmt.Println("Нажмите любую кнопку для продолжения")
		bufio.NewScanner(os.Stdin).Scan()

		matrix = makeStep(fieldSize, matrix)
	}
}
```

14. **Цифровой корень числа** 

Напишите программу, которая вычисляет цифровой корень числа. Цифровой корень — это рекурсивная сумма цифр числа, пока не останется только одна цифра. Например, цифровой корень числа 9875 равен 2, потому что 
9+8+7+5=29 → 2+9=11 → 1+1=2

```go
package main

import (
	"fmt"
)

func digitsSqrt(number int) int {

	currentNum := 0
	for ok := true; ok; ok = (number/10 != 0) {
		for ; number > 0; number /= 10 {
			currentNum += number % 10
		}

		number = currentNum
		currentNum = 0
	}

	return number
}

func main() {

	var num int
	fmt.Println("Введите число:")
	fmt.Scan(&num)

	fmt.Print(digitsSqrt(num))
}

```

15. **Римские цифры**
    
Напишите функцию, которая преобразует арабское число (например, 1994) в римское (например, \"MCMXCIV\").
Программа должна использовать циклы и условные операторы для создания римской записи.
```go
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

```

      
