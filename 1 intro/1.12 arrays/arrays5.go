package main

import "fmt"

func main() {
	var n, sum int
	fmt.Scan(&n)
	arr := make([]int, n, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	for _, elem := range arr {
		if elem > 0 {
			sum++
		}
	}
	fmt.Print(sum)
}

// Дана последовательность, состоящая из целых чисел. Напишите программу,
// которая подсчитывает количество положительных чисел среди элементов последовательности.
