package main

import "fmt"

func main() {
	var e, max, sum int

	for fmt.Scan(&e); e != 0; fmt.Scan(&e) {
		if e == max {
			sum++
		} else if e > max {
			max = e
			sum = 1
		}
	}
	fmt.Print(sum)
}

// Последовательность состоит из натуральных чисел и завершается числом 0.
// Определите количество элементов этой последовательности, которые равны ее наибольшему элементу.
