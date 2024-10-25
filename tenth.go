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
