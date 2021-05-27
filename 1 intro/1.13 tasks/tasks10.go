package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	max := (y / 7) * 7
	if y < 0 {
		fmt.Print(max - 7)
	} else if max >= x {
		fmt.Print(max)
	} else {
		fmt.Print("NO")
	}
}

// Найдите самое большее число на отрезке от a до b, кратное 7.
