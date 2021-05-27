package main

import "fmt"

func main() {
	var n int
	var a int = 1

	fmt.Scan(&n)

	for a <= n {
		fmt.Print(a, " ")
		a *= 2
	}
}

// По данному числу N распечатайте все целые значения степени двойки, не превосходящие N, в порядке возрастания.
