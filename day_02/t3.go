// 3. Дано неотрицательное целое число. Найдите и выведите первую цифру числа.
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int

	fmt.Scan(&n)

	nStr := strconv.Itoa(n)

	fmt.Println(string(nStr[0]))
}
