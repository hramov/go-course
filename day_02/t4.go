// 4. Определите является ли билет счастливым. Счастливым считается билет, в шестизначном номере которого сумма первых трёх цифр совпадает с суммой трёх последних.
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int

	fmt.Scan(&n)

	nStr := strconv.Itoa(n)

	fmt.Println(nStr[0]+nStr[1]+nStr[2] == nStr[3]+nStr[4]+nStr[5])
}
