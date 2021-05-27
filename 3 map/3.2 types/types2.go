func adding(first, second string) int64 {
	firstNumber, _ := strconv.ParseInt(cleaner(first), 10, 64)
	secondNumber, _ := strconv.ParseInt(cleaner(second), 10, 64)

	return firstNumber + secondNumber
}

func cleaner(s string) string {
	tmp := make([]rune, 0)

	for _, r := range []rune(s) {
		if unicode.IsDigit(r) {
			tmp = append(tmp, r)
		}
	}

	return string(tmp)
}

// Представьте что вы работаете в большой компании где используется модульная архитектура.
// Ваш коллега написал модуль с какой-то логикой (вы не знаете) и передает в вашу программу какие-то данные.
// Вы же пишете функцию которая считывает две переменных типа string,
// а возвращает число типа int64 которое нужно получить сложением этих строк.
