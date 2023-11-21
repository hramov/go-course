// 5. Напишите программу, которая в последовательности чисел находит сумму двузначных чисел, кратных 8.
// Программа в первой строке получает на вход число n чисел, входящих в данную последовательность.
// Input: 38 24 800 8 16, Output: 40
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var sum int

	in := bufio.NewReader(os.Stdin)

	line, err := in.ReadString('\n')
	if err != nil {
		panic(err)
	}

	line = strings.TrimSuffix(line, "\n")
	line = strings.TrimSuffix(line, "\r")

	arr := strings.Split(line, " ")

	for _, v := range arr {
		i, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		if i%8 == 0 && i > 9 && i < 100 {
			sum += i
		}
	}

	fmt.Println(sum)
}
