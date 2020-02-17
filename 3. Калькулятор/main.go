package main

import "fmt"

func main() {
	var a, b int
	var operator string

	fmt.Scan(&a, &b, &operator)

	if operator == "+" {
		println(a + b)
	} else if operator == "-" {
		println(a - b)
	} else if operator == "*" {
		println(a * b)
	} else if operator == "/" {
		println(a / b)
	} else {
		println("Неверный знак")
	}
}
