func merge2Channels(f func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	results := make([]int, n)
	go func() {
		wg := new(sync.WaitGroup)
		wg.Add(n)
		for i := 0; i < n; i++ {
			x1 := <-in1
			x2 := <-in2
			index := i
			go func() {
				x1res := make(chan int)
				x2res := make(chan int)
				go func() {
					x1res <- f(x1)
				}()
				go func() {
					x2res <- f(x2)
				}()
				results[index] = (<-x1res) + (<-x2res)
				wg.Done()
			}()
		}
		wg.Wait()
		for i := 0; i < n; i++ {
			out <- results[i]
		}
	}()
}

// Необходимо написать функцию func merge2Channels(fn func(int) int, in1 <-chan int, in2 <- chan int, out chan<- int, n int).
// Описание ее работы:
// n раз сделать следующее
// прочитать по одному числу из каждого из двух каналов in1 и in2, назовем их x1 и x2.
// вычислить f(x1) + f(x2)
// записать полученное значение в out
// Функция merge2Channels должна быть неблокирующей, сразу возвращая управление.
// Функция fn может работать долгое время, ожидая чего-либо или производя вычисления.
