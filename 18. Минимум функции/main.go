package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, x, y float64

	fmt.Scan(&a, &b, &c)

	if a == 0 {
		println("а равняется нулю")
		return
	}

	x = minimumX(a, b)
	y = minimumY(a, b, c, x)

	println(x, y)
}

func minimumX(a, b float64) float64 {
	return -b / (2 * a)
}

func minimumY(a, b, c, x float64) float64 {
	return a*math.Pow(x, 2) + b*x + c
}
