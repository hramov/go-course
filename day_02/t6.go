// 6. Вклад в банке составляет x рублей. Ежегодно он увеличивается на p процентов, после чего дробная часть копеек отбрасывается.
// Каждый год сумма вклада становится больше.
// Определите, через сколько лет вклад составит не менее y рублей.

package main

import (
	"fmt"
	"math"
)

func roundFloat(val float32, precision uint) float32 {
	ratio := math.Pow(10, float64(precision))
	return float32(math.Round(float64(val)*ratio) / ratio)
}

func main() {
	var x float32
	var p float32
	var y float32

	fmt.Scan(&x)
	fmt.Scan(&p)
	fmt.Scan(&y)

	c := 0

	for x < y {
		x = roundFloat(x+x*p/100, 2)
		c++
	}

	fmt.Println(c)
}
