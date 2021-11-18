package main

import (
	"strconv"
	"strings"
)

func Add(numbers string) int {
	result := 0

	numberArray := strings.Split(numbers, ",")

	for _, value := range numberArray {
		convertedValue, _ := strconv.Atoi(value)
		result += convertedValue
	}

	return result
}
