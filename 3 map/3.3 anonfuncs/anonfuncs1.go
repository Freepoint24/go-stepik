fn := func(i uint) uint {
	iStr := strconv.FormatUint(uint64(i), 10)
	i = 0
	for _, ch := range iStr {
		temp, _ := strconv.ParseUint(string(ch), 10, 32)
		parsed := uint(temp)
		if parsed%2 == 0 && parsed != 0 {
			i *= 10
			i += parsed
		}
	}
	if i == 0 {
		i = 100
	}
	return i
}

// Функция на вход получает целое положительное число (uint), возвращает число того же типа,
// в котором отсутствуют нечетные цифры и цифра 0. Если же получившееся число равно 0, то возвращается число 100.
// Цифры в возвращаемом числе имеют тот же порядок, что и в исходном числе.
