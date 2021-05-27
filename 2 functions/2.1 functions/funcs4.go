func fibonacci(n int) int {
	f1, f2 := 0, 1
	if n == 0 {
		return 0
	} else {
		for i := 1; i < n; i++ {
			f1, f2 = f2, f1+f2
		}
	}
	return f2
}

// Последовательность Фибоначчи определена следующим образом: φ1=1, φ2=1, φn=φn-1+φn-2 при n>1.
// Начало ряда Фибоначчи выглядит следующим образом: 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, ...
// Напишите функцию, которая по указанному натуральному n возвращает φn.