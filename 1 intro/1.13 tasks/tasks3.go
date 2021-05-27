package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	h := n / 3600
	m := n % 3600 / 60
	fmt.Printf("It is %d hours %d minutes.", h, m)
}

// Идёт k-я секунда суток. Определите, сколько целых часов h и целых минут m прошло с начала суток. Например, если
// k=13257=3*3600+40*60+57, то h=3 и m=40.
