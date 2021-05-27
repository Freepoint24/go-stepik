func calculator(arguments <-chan int, done <-chan struct{}) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		var sum, value int
		for {
			select {
			case value = <-arguments:
				sum += value
			case <-done:
				out <- sum
				return
			}
		}
	}()
	return out
}

// В качестве аргумента эта функция получает два канала только для чтения, возвращает канал только для чтения.
// Через канал arguments функция получит ряд чисел, а через канал done - сигнал о необходимости завершить работу.
// Когда сигнал о завершении работы будет получен, функция должна в выходной (возвращенный) канал отправить сумму полученных чисел.
// Функция calculator должна быть неблокирующей, сразу возвращая управление.
