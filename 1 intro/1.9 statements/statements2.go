package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	x := a / 100
	y := a / 10 % 10
	z := a % 10
	if x != y && x != z && y != z {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}

// По данному трехзначному числу определите, все ли его цифры различны.
