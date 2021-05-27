package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	temp := n % 10
	switch {
	case n >= 11 && n <= 14:
		fmt.Print(n, " korov")
	case temp == 0 || (temp >= 5 && temp <= 9):
		fmt.Print(n, " korov")
	case temp == 1:
		fmt.Print(n, " korova")
	case temp >= 2 && temp <= 4:
		fmt.Print(n, " korovy")
	}
}

// По данному числу n закончите фразу "На лугу пасется..."
// одним из возможных продолжений: "n коров", "n корова", "n коровы", правильно склоняя слово "корова".
