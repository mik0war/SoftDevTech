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
