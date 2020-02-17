package main

import (
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	num := rand.Intn(900) + 100
	println("Число", num)

	c := num % 10
	b := (num%100 - c) / 10
	a := (num - c - b) / 100
	println("Сумма цифр", a+b+c)
	println("Произведение цифр", a*b*c)
}
