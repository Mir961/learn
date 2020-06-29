package main

import (
	"fmt"
	"math"
)

func main() {
	var circleX, circleY, circleRadius, pointX, pointY float64

	println("Введите кординаты окружности")
	fmt.Scan(&circleX, &circleY)
	println("Введите радиус")
	fmt.Scan(&circleRadius)
	println("Введите кординаты точки")
	fmt.Scan(&pointX, &pointY)

	distance := getDistance(circleX, circleY, pointX, pointY)

	println(getAnswer(distance, circleRadius))
}

func getDistance(x1, y1, x2, y2 float64) float64 {
	// Возвращает расстояние между двумя точками
	a := x1 - x2
	b := y1 - y2

	return math.Pow((a*a)+(b*b), 0.5)
}

func getAnswer(distance, circleRadius float64) string {
	if distance > circleRadius {
		return "Точка вне окружностью"
	}

	if distance < circleRadius {
		return "Точка внутри окружности"
	}

	return "Точка на границе окружности"
}
