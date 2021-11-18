package main

import (
	"fmt"
	"strings"
)

func Add(numbers string) int {
	result := 0

	numberArray := strings.Split(numbers, ",")

	for _, value := range numberArray {
		fmt.Printf("%s", value)
		fmt.Printf("--------------")
		//result += fmt.Sscan(value)
	}

	return result
}
