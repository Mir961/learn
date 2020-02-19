package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64

	fmt.Scan(&a, &b, &c)

	if discriminant(a, b, c) < 0 {
		println("Данное уравнение не имеет корней")
	} else {
		x1, x2 := roots(a, b, discriminant(a, b, c))
		println("Корни", x1, x2)
	}
}

func discriminant(a float64, b float64, c float64) float64 {

	dis := (b * b) - (4 * a * c)
	return dis
}

func roots(a float64, b float64, discriminant float64) (float64, float64) {

	x1 := (-b + math.Pow(discriminant, 0.5)) / (2 * a)
	x2 := (-b - math.Pow(discriminant, 0.5)) / (2 * a)
	return x1, x2

}
