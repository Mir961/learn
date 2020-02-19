package main

import "fmt"

func main() {
	var a, b, c int

	println("Введите три числа")
	fmt.Scan(&a, &b, &c)
	println(average(a, b, c))
}

func average(a, b, c int) float64 {
	return (float64(a + b + c)) / 3
}
