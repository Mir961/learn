package main

import "fmt"

func main() {
	var input string

	fmt.Scan(&input)
	inputRunes := []rune(input)
	length := len(inputRunes)

	fmt.Println(string(inputRunes[:length/2]))
	fmt.Println(string(inputRunes[length/2:]))
}
