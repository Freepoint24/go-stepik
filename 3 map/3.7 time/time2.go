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
	t, _ := time.Parse("2006-01-02 15:04:05", timeStr)
	if t.Hour() > 13 {
		t = t.AddDate(0, 0, 1)
	}
	fmt.Print(t.Format("2006-01-02 15:04:05"))
}

// На стандартный ввод подается строковое представление даты и времени определенного события в следующем формате:
// 2020-05-15 08:00:00
// Если время события до обеда (13-00), то ничего менять не нужно, достаточно вывести дату на стандартный вывод в том же формате.
// Если же событие должно произойти после обеда, необходимо перенести его на то же время на следующий день,
// а затем вывести на стандартный вывод в том же формате.