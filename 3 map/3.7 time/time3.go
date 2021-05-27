package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	timeStr, _ := rd.ReadString('\n')
	timeStr = strings.TrimSuffix(timeStr, "\n")
	tArr := strings.Split(timeStr, ",")
	date1, _ := time.Parse("02.01.2006 15:04:05", tArr[0])
	date2, _ := time.Parse("02.01.2006 15:04:05", tArr[1])
	if date1.Before(date2) {
		fmt.Println(date2.Sub(date1))
	} else {
		fmt.Println(date1.Sub(date2))
	}
}

// На стандартный ввод подается строковое представление двух дат, разделенных запятой (формат данных смотрите в примере).
// Необходимо преобразовать полученные данные в тип Time, а затем вывести продолжительность периода между меньшей и большей датами.
