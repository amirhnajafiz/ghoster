package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args[1:]) < 2 {
		panic("not enough arguments")
	}

	parts := strings.Split(os.Args[1], os.Args[2])

	for _, part := range parts {
		fmt.Println(part)
	}
}
