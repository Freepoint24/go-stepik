package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a string
	fmt.Scan(&a)
	for _, ch := range a {
		m, _ := strconv.Atoi(string(ch))
		fmt.Print(m * m)
	}
}

// На вход подается целое число. Необходимо возвести в квадрат каждую цифру числа и вывести получившееся число.
