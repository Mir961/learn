package main

import (
	"fmt"
)

func main() {
	var panels, a, b int

	println("Условие: панели (1 ≤ N ≤ 100), A (1 ≤ A ≤ 100), B (1 ≤ B ≤ 100).")
	fmt.Scan(&panels, &a, &b)

	println("Количество нанограм сульфида", panels*a*b*2)
}
