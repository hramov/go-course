// 3. Напишите функцию, которая умножает значения на которые ссылаются два указателя и выводит получившееся произведение в консоль.

package main

import "fmt"

func pointer(p1, p2 *int) {
	fmt.Println(*p1 * *p2)
}

func main() {
	p1 := 6
	p2 := 2

	pointer(&p1, &p2)
}
