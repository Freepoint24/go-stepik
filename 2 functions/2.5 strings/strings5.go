package main

import (
	"fmt"
	"strings"
)

func main() {
	var n string
	fmt.Scan(&n)
	arr := []rune(n)
	for _, e := range arr {
		if strings.Count(n, string(e)) > 1 {
			n = strings.ReplaceAll(n, string(e), "")
		}
	}
	fmt.Print(n)
}

// Дается строка. Нужно удалить все символы, которые встречаются более одного раза и вывести получившуюся строку.
