package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, dis float64

	fmt.Scan(&a, &b, &c)
	dis = discriminant(a, b, c)

	if dis < 0 {
		println("Данное уравнение не имеет корней")
		return
	}

	x1, x2 := roots(a, b, dis)
	println("Корни", x1, x2)
}

func discriminant(a, b, c float64) float64 {
	return (b * b) - (4 * a * c)
}

func roots(a, b, dis float64) (float64, float64) {
	x1 := (-b + math.Pow(dis, 0.5)) / (2 * a)
	x2 := (-b - math.Pow(dis, 0.5)) / (2 * a)

	return x1, x2
}
