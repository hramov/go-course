// 5. Создать функцию divide которая делит два числа, но возвращает не только результат деления, но и информацию об ошибке.
// В случае какой-либо ошибки нужно вывести "ошибка", если ошибки нет - результат функции.
// Функция divide(a int, b int) (int, error) получает на вход два числа которые нужно поделить и возвращает результат (int) и ошибку (error)

package main

import "fmt"

func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func main() {
	fmt.Println(divide(5, 2))
	fmt.Println(divide(1, 0))
}
