package main

import (
	"fmt"
	"strings"
)

func main() {
	var a string
	fmt.Scan(&a)
	fmt.Println(strings.Join(strings.Split(a, ""), "*"))
}

// Дана строка, содержащая только английские буквы (большие и маленькие).
// Добавить символ ‘*’ (звездочка) между буквами (перед первой буквой и после последней символ ‘*’ добавлять не нужно).
