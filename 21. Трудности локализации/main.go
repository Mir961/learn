package main

import "fmt"

func main() {
	var a int

	println("Введите колличество монстров")
	fmt.Scan(&a)

	println(quantity(a))
}

func quantity(a int) string {
	if a >= 1 && a <= 4 {
		return "few"
	}

	if a >= 5 && a <= 9 {
		return "several"
	}

	if a >= 10 && a <= 19 {
		return "pack"
	}

	if a >= 20 && a <= 49 {
		return "lots"
	}

	if a >= 50 && a <= 99 {
		return "horde"
	}

	if a >= 100 && a <= 249 {
		return "throng"
	}

	if a >= 250 && a <= 499 {
		return "swarm"
	}

	if a >= 500 && a <= 999 {
		return "zounds"
	}

	if a >= 1000 {
		return "legion"
	}
	panic("Невозможно определить слово")
}
