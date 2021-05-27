package main

import "fmt"

func main() {
	var a string
	var max int32
	fmt.Scan(&a)
	for _, ch := range a {
		if ch > max {
			max = ch
		}
	}
	fmt.Print(string(max))
}

// Дана строка, содержащая только десятичные цифры. Найти и вывести наибольшую цифру.
// Вводится строка ненулевой длины. Известно также, что длина строки не превышает 1000 знаков
// и строка содержит только десятичные цифры.
