func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	res := make(chan int, 1)
	go func() {
		defer close(res)
		select {
		case f := <-firstChan:
			res <- f * f
		case s := <-secondChan:
			res <- s * 3
		case <-stopChan:
			return
		}
	}()
	return res
}

// Функция получает в качестве аргументов 3 канала, и возвращает канал типа <-chan int.
// в случае, если аргумент будет получен из канала firstChan, в выходной (возвращенный)
// канал вы должны отправить квадрат аргумента.
// в случае, если аргумент будет получен из канала secondChan, в выходной (возвращенный)
// канал вы должны отправить результат умножения аргумента на 3.
// в случае, если аргумент будет получен из канала stopChan, нужно просто завершить работу функции.
