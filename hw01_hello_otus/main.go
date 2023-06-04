package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	input := "Hello, OTUS!"
	fmt.Println(stringutil.Reverse(input))
}
