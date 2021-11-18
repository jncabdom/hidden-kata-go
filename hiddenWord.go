package main

import (
	"fmt"
)

var translation = map[rune]rune{
	'0': 'o',
	'1': 'b',
	'2': 'l',
	'3': 'i',
	'4': 'e',
	'5': 't',
	'6': 'a',
	'7': 'd',
	'8': 'n',
	'9': 'm',
}

func HiddenWord(code uint32) string {
	codeString := fmt.Sprint(code)
	result := ""
	for _, current := range codeString {
		result += string(translation[current])
	}
	return result
}
