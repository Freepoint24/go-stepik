package main

import (
	"fmt"
	"regexp"
)

func main() {
	var n string
	fmt.Scan(&n)

	var IsPass = regexp.MustCompile(`^^[a-zA-Z0-9]{5,}$`).MatchString

	if IsPass(n) {
		fmt.Print("Ok")
	} else {
		fmt.Print("Wrong password")
	}
}

// Ваша задача сделать проверку подходит ли пароль вводимый пользователем под заданные требования.
// Длина пароля должна быть не менее 5 символов, он должен содержать только цифры и буквы латинского алфавита.
// На вход подается строка-пароль. Если пароль соответствует требованиям - вывести "Ok", иначе вывести "Wrong password".
