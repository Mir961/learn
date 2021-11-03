package main

import (
	"fmt"
)

func main() {
	var deposit, rate, term, total, difference float64

	fmt.Scan(&deposit, &rate, &term)

	rate = rate / 100 / 12
	total = depositAfter(deposit, rate, term)
	difference = depositDifference(total, deposit)

	println(total, difference)
}

func depositAfter(deposit, rate, term float64) float64 {
	var i, total float64

	total = deposit

	for i = 0; i <= term; i++ {
		total = total + total*rate
	}
	return total
}

func depositDifference(total, deposit float64) float64 {
	return total - deposit
}
