package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const now = 1589570165

func main() {
	rd := bufio.NewReader(os.Stdin)
	timeStr, _ := rd.ReadString('\n')
	timeStr = strings.TrimSuffix(timeStr, "\n")
	timeStr = strings.ReplaceAll(timeStr, " мин. ", "m")
	timeStr = strings.ReplaceAll(timeStr, " сек.", "s")
	dur, _ := time.ParseDuration(timeStr)
	timein := time.Unix(now, 0).Add(dur).Format("Mon Jan _2 15:04:05 UTC 2006")
	fmt.Println(timein)
}

// На стандартный ввод подаются данные о продолжительности периода (формат приведен в примере).
// Кроме того, вам дана дата в формате Unix-Time: 1589570165 в виде константы типа int64.
// Требуется считать данные о продолжительности периода, преобразовать их в тип Duration,
// а затем вывести (в формате UnixDate) дату и время, получившиеся при добавлении периода к стандартной дате.
