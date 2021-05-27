package main

import "fmt"

func main() {
	var x, p, y float32
	var i int

	fmt.Scan(&x, &p, &y)
	for i = 0; x <= y; i++ {
		x += x * (p / 100)
	}
	fmt.Println(i)
}

// Вклад в банке составляет x рублей.
// Ежегодно он увеличивается на p процентов, после чего дробная часть копеек отбрасывается.
// Каждый год сумма вклада становится больше. Определите, через сколько лет вклад составит не менее y рублей.