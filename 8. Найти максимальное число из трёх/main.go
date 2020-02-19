package main

import "fmt"

func main() {
	var a, b, c, max int

	println("Введите три разных числа")
	fmt.Scan(&a, &b, &c)

	if a >= b {
		max = a
	} else {
		max = b
	}

	if c > max {
		max = c
	}

	println("Максимальное из них", max)
}
