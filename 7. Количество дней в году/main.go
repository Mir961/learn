package main

import "fmt"

func main() {
	var year int

	println("Введите год")
	fmt.Scan(&year)

	if isLeapYear(year) {
		println("Год високосный")
		return
	}

	println("Год не високосный")
}

func isLeapYear(year int) bool {
	return (year%400 == 0) || (year%4 == 0) && (year%100 != 0)
}
