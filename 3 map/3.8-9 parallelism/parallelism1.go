func task(i chan int, n int) {
	i <- n + 1
}

// Напишите функцию которая принимает канал и число N типа int. Необходимо вернуть значение N+1 в канал. 