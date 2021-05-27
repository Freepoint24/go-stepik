package main

import "fmt"

func main() {
	var number, digit int
	fmt.Scan(&number, &digit)
	result := 0
	for i := 1; number > 0; {
		if number%10 != digit {
			result += number % 10 * i
			i *= 10
		}
		number /= 10
	}

	fmt.Println(result)
}

// Из натурального числа удалить заданную цифру.
