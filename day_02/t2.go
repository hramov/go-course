// 2. По введенному трехзначному числу определите, все ли его цифры различны.
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int

	fmt.Scan(&n)

	m := make(map[rune]struct{})

	for _, v := range strconv.Itoa(n) {
		if _, ok := m[v]; ok {
			fmt.Println("Совпадение")
			return
		}
		m[v] = struct{}{}
	}

	fmt.Println("Разные")
}
