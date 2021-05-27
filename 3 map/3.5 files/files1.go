func main() {
	var res int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		n, _ := strconv.Atoi(s)
		res += n
	}
	io.WriteString(os.Stdout, strconv.Itoa(res))
}

// Задача состоит в следующем: на стандартный ввод подаются целые числа в диапазоне 0-100,
// каждое число подается на стандартный ввод с новой строки (количество чисел не известно).
// Требуется прочитать все эти числа и вывести в стандартный вывод их сумму.
