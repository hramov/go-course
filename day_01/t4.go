package main

import "fmt"

func main() {
	var angle int

	fmt.Scan(&angle)

	hours := angle * 2 / 60

	minutes := angle * 2 % 60

	fmt.Printf("It's %d hours and %d minutes", hours, minutes)
}
