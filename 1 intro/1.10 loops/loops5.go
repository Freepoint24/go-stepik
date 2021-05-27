package main

import "fmt"

func main() {
	var n, c, d int
	fmt.Scan(&n, &c, &d)

	for i := 0; i <= n; i++ {
		if (i%c == 0) && (i%d != 0) {
			fmt.Print(i)
			break
		}
	}
}

// Найдите первое число от 1 до n включительно, кратное c, но НЕ кратное d.
