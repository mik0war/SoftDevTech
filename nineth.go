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
