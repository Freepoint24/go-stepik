package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Print(arr[3])
}

// Напишите программу, принимающая на вход число N(N>=4), а затем NN целых чисел-элементов среза.
// На вывод нужно подать 4-ый (3 по индексу) элемент данного среза.
