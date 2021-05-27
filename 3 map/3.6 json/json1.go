package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type (
	Student struct {
		Rating []int
	}
	Group struct {
		Students []Student
	}
)

func main() {
	students := new(Group)
	bytes, _ := ioutil.ReadAll(os.Stdin)
	_ = json.Unmarshal(bytes, students)
	var cnt int

	for _, student := range students.Students {
		cnt += len(student.Rating)
	}

	average := float64(cnt) / float64(len(students.Students))
	bytes, _ = json.MarshalIndent(map[string]float64{"Average": average}, "", "    ")
	fmt.Printf("%s", bytes)
}

// В сведениях о каждом студенте содержится информация о полученных им оценках (Rating).
// Требуется прочитать данные, и рассчитать среднее количество оценок, полученное студентами группы.
