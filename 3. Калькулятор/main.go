package main

import "fmt"

func main() {
	var a, b int
	var operation string

	_, err := fmt.Scan(&a, &operation, &b)

	if err != nil {
		println("Введены некорректные данные")
		return
	}

	println(Calculate(a, operation, b))
}

func Calculate(a int, operation string, b int) int {
	if operation == "+" {
		return a + b
	}

	if operation == "-" {
		return a - b
	}

	if operation == "*" {
		return a * b
	}

	if operation == "/" {
		return a / b
	}

	panic("Невозможно определить операцию")
}
