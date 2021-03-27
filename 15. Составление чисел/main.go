package main

import "fmt"

func main() {
	var num int

	println("Введите трехзначное число")
	fmt.Scan(&num)

	println(makingNumbers(num))
}

func makingNumbers(num int) int {

	c := num % 10
	b := (num % 100) / 10
	a := num / 100
	println(a*100 + c*10 + b)
	println(b*100 + a*10 + c)
	println(b*100 + c*10 + a)
	println(c*100 + a*10 + b)

	return (c*100 + b*10 + a)
}
