package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
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
	if len(os.Args[1:]) < 1 {
		panic("not enough arguments")
	}

	limit := convertToIntOrPanic(os.Args[1])

	time.Sleep(time.Duration(limit) * time.Second)

	fmt.Println("finished!")
}
