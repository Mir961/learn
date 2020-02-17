package main

import (
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var a int

	a = rand.Intn(900)

	println(a + 100)
}
