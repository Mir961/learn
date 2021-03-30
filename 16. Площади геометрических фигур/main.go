package main

import "fmt"

func main() {
	var i, a, b, c, h float64

	fmt.Println("Введите номер фигуры (1 - квадрат, 2 - прямоугольник, 3 - треугольник, 4 - круг)")
	fmt.Scan(&i)

	switch i {
	case 1:
		fmt.Println(square(), "Площадь квадрата")
	case 2:
		fmt.Println(rectangle(), "Площадь прямоугольника")
	case 3:
		println("Введите три стороны и высоту к первой стороне")
		fmt.Scan(&a, &b, &c, &h)

		if existenceTriangle(a, b, c) {
			fmt.Println(triangle(a, h), "Площадь треугольника")
		} else {
			fmt.Println("Треугольник не существует")
		}
	case 4:
		fmt.Println(circle(), "Площадь круга")
	default:
		fmt.Println("Неизвестный номер фигуры")
	}
}

func triangle(a, h float64) float64 {
	return 0.5 * a * h
}

func existenceTriangle(b, c, a float64) bool {
	//проверяет существование треугольника
	return (a < b+c) && (b < a+c) && (c < a+b)
}

func square() int {
	var a int

	fmt.Println("Введите сторону квадрата")
	fmt.Scan(&a)

	return a * a
}

func rectangle() int {
	var a, b int

	fmt.Println("Введите стороны прямоугольника")
	fmt.Scan(&a, &b)

	return a * b
}

func circle() float64 {
	var r float64

	fmt.Println("Введите радиус круга")
	fmt.Scan(&r)

	return 0.5 * r * r
}
