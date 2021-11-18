package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringCalculator(t *testing.T) {
	t.Run("Returns 0 when getting an empty string", func(t *testing.T) {
		actual := Add("")
		expected := 0

		assert.Equal(t, expected, actual)
	})

	t.Run("Returns the value if only got one", func(t *testing.T) {
		actual := Add("1")
		expected := 1

		assert.Equal(t, expected, actual)
	})

	t.Run("Returns the addition of both numbers", func(t *testing.T) {
		actual := Add("1,2")
		expected := 3

		assert.Equal(t, expected, actual)
	})

	t.Run("It can add an unknown amount of numbers", func(t *testing.T) {
		actual := Add("1,2,3,4,5,6")
		expected := 21

		assert.Equal(t, expected, actual)
	})
}
