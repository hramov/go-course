package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var c int

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
		if i > 0 {
			c++
		}
	}

	fmt.Println(c)
}
