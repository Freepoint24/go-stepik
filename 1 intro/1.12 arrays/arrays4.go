package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	for indx, elem := range arr {
		if indx%2 == 0 {
			fmt.Print(elem, " ")
		}
	}
}

// Дан массив, состоящий из целых чисел. Нумерация элементов начинается с 0.
// Напишите программу, которая выведет элементы массива, индексы которых четны (0, 2, 4...).
