package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	n *= 2
	n += 100

	fmt.Println(n)
}
