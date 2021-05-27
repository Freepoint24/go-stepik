package main

import "fmt"

func main() {
	var n, cur, min, count int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&cur)
		if cur < min || i == 0 {
			count = 1
			min = cur
		} else if cur == min {
			count++
		}
	}
	fmt.Print(count)
}

// Найдите количество минимальных элементов в последовательности.
