package main

import "fmt"

func main() {
	var n, x, sum int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		if x == 0 {
			sum++
		}
	}
	fmt.Print(sum)
}

// По данным числам, определите количество чисел, которые равны нулю.
