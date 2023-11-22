// 2. Напишите функцию sumInt(), принимающую переменное количество аргументов типа int, и возвращающую количество полученных функцией аргументов и их сумму.

package main

import "fmt"

func sumInt(n ...int) (int, int) {
	sum := 0

	for _, v := range n {
		sum += v
	}

	return len(n), sum
}

func main() {
	fmt.Println(sumInt(1, 2, 3))
}
