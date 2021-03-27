package main

import "fmt"

func main() {
	var a, b, c int

	println("Введите длины сторон треугольника")
	fmt.Scan(&a, &b, &c)

	if exTriangle(a, b, c) {
		println("Треугольник существует")
		return
	}

	println("Треугольник не существует")
}

func exTriangle(a, b, c int) bool {
	return (a < b+c) && (b < a+c) && (c < a+b)
}
