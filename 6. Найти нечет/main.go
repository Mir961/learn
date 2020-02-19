package main

import "fmt"

func main() {
	var a, b int

	println("Введите два числа с разной четностью")
	fmt.Scan(&a, &b)

	if a%2 != 0 {
		println("Нечетное число", a)
		return
	}

	println("Нечетное число", b)
}
