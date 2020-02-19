package main

import (
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	num := (rand.Intn(90) + 10) * 10
	println("Число", num)
}
