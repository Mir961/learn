package main

import "fmt"

func main() {
	var a, b int

	println("Введите два числа")
	fmt.Scan(&a, &b)

	if a%b == 0 {
		println("Первое число кратно второму")
		return
	}

	println("Первое число не кратно второму")
}
