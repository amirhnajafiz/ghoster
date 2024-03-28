package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func convertToIntOrPanic(input string) int {
	tmp, err := strconv.Atoi(input)
	if err != nil {
		log.Println(err)

		panic("input argument is not an integer type")
	}

	return tmp
}

func main() {
	if len(os.Args[1:]) < 2 {
		panic("not enough arguments")
	}

	a := convertToIntOrPanic(os.Args[1])
	b := convertToIntOrPanic(os.Args[2])

	fmt.Println(a + b)
}
