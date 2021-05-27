package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)

	sum := 0
	for ; a <= b; {
		sum += a
		a++
	}
	fmt.Println(sum)
}

// Требуется написать программу, при выполнении которой с клавиатуры считываются два натуральных числа A и B
// (каждое не более 100, A < B). Вывести сумму всех чисел от A до B  включительно.
