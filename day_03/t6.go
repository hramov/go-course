// 6. Валидация пароля: Длина пароля не менее 5 символов, должен содержать только арабские цифры и буквы латинского алфавита.
// На вход подается строка-пароль.
// Если пароль соответствует требованиям - вывести "Ok", иначе вывести "Wrong password
package main

import (
	"fmt"
)

var allowed = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', '1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func contains(r rune) bool {
	for _, v := range allowed {
		if v == r {
			return true
		}
	}
	return false
}

func isMoreOrEqualThan5Symbols(p string) bool {
	return len(p) >= 5
}

func isValidFormat(p string) bool {
	for _, v := range []rune(p) {
		if !contains(v) {
			return false
		}
	}
	return true
}

func validate(p string) bool {
	return isMoreOrEqualThan5Symbols(p) && isValidFormat(p)
}

func main() {
	var p string

	fmt.Scan(&p)

	if validate(p) {
		fmt.Println("Ok")
	} else {
		fmt.Println("Wrong password")
	}
}
