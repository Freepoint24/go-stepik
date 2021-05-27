package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	switch {
	case a+b > c && a+c > b && b+c > a:
		fmt.Print("Существует")
	default:
		fmt.Print("Не существует")
	}
}

// Даны три натуральных числа a, b, c. Определите, существует ли треугольник с такими сторонами.
