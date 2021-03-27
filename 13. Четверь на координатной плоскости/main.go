package main

import "fmt"

func main() {
	var x, y int

	println("Введите координаты точки")
	fmt.Scan(&x, &y)

	println(quarterNumber(x, y), "четверть")
}

func quarterNumber(x int, y int) int {
	if x >= 0 && y >= 0 {
		return 1
	}

	if x <= 0 && y >= 0 {
		return 2
	}

	if x <= 0 && y <= 0 {
		return 3
	}

	if x >= 0 && y <= 0 {
		return 4
	}
	panic("Невозможно определить четверть")
}
