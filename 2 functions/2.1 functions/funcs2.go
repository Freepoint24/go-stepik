func minimumFromFour() int {
	var n, nmin int
	for i := 1; i < 5; i++ {
		fmt.Scan(&n)
		if i == 1 {
			nmin = n
		}
		if i > 1 && n < nmin {
			nmin = n
		}
	}
	return nmin
}

// Напишите функцию, находящую наименьшее из четырех введённых в этой же функции чисел.
