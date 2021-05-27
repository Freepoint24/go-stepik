package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	switch {
	case 1 <= a && a <= 9:
		fmt.Print(a)
	case 10 <= a && a <= 99:
		fmt.Print(a / 10)
	case 100 <= a && a <= 999:
		fmt.Print(a / 100)
	case 1000 <= a && a <= 9999:
		fmt.Print(a / 1000)
	case 10000 <= a && a <= 99999:
		fmt.Print(a / 10000)
	}
}

// Дано неотрицательное целое число. Найдите и выведите первую цифру числа.
