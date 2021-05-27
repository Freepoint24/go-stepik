package main

import "fmt"

func main() {
	var n string
	fmt.Scan(&n)
	rn := []rune(n)

	for i, e := range rn {
		if i%2 != 0 {
			fmt.Print(string(e))
		}
	}
}

// На вход дается строка, из нее нужно сделать другую строку, оставив только нечетные символы (считая с нуля).
