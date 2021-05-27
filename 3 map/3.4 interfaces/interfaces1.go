package main

import (
	// пакет используется для проверки ответа, не удаляйте его
	"fmt" // пакет используется для проверки ответа, не удаляйте его
)

func main() {
	value1, value2, operation := readTask()
	vi := [2]interface{}{value1, value2}
	var v1, v2 float64
	vf := [2]*float64{&v1, &v2}
	var ok bool
	for i, v := range vi {
		if *vf[i], ok = v.(float64); !ok {
			fmt.Printf("value=%v: %T", v, v)
			return
		}
	}
	ops := map[string]func() float64{
		"+": func() float64 { return v1 + v2 },
		"-": func() float64 { return v1 - v2 },
		"*": func() float64 { return v1 * v2 },
		"/": func() float64 { return v1 / v2 },
	}

	if oper, ok := operation.(string); ok {
		if fun, ok := ops[oper]; ok {
			fmt.Printf("%.4f", fun())
			return
		}
	}
	fmt.Print("неизвестная операция")

}

// Итак, вы получаете 3 значения типа пустой интерфейс: если все удачно, то первые 2 значения будут float64.
// Третье значение в идеальном случае будет строкой, которая может иметь значения: "+", "-", "*", "/"
// (определенная математическая операция). Но такие идеальные случаи будут не всегда, вы можете получить значения других типов,
// а также строка в третьем значении может не относится к одной из указанных математических операций.
