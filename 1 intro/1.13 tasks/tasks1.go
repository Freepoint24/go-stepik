package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Print(n/100 + n/10%10 + n%10)
}

// Дано трехзначное число. Найдите сумму его цифр.
