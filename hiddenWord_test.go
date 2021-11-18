package main

import "testing"

func TestHiddenWord(t *testing.T) {
	const example1 = 147
	const result1 = "bed"

	if HiddenWord(example1) != result1 {
		t.Errorf("hiddenWord(%d) => %s", example1, HiddenWord(example1))
	}
}
