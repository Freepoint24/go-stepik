package main

import (
	"fmt"
)

func main() {
	var (
		first, second int
		reverse       int
	)

	fmt.Scan(&first, &second)

	for first > 0 {
		first, reverse = first/10, reverse*10+first%10
	}

	for reverse > 0 {
		current := reverse % 10
		reverse /= 10
		tmp := second

		for tmp > 0 {
			if current == tmp%10 {
				fmt.Printf("%d ", current)
			}
			tmp /= 10
		}
	}
}

// Даны два числа. Определить цифры, входящие в запись как первого, так и второго числа.
