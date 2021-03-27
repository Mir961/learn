package main

import "fmt"

func main() {
	var i int

	fmt.Println("Введите номер фигуры")
	fmt.Scan(&i)

	switch i {
	case 1:
		fmt.Println("Квадрат")
	case 2:
		fmt.Println("Прямоугольник")
	case 3:
		fmt.Println("Треуголник")
	case 4:
		fmt.Println("Круг")
	default:
		fmt.Println("Неизвестный номер")
	}
}
