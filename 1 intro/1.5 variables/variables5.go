package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	fmt.Print(a / 10 % 10)
}

// Дано неотрицательное целое число. Найдите число десятков (то есть вторую цифру справа).
