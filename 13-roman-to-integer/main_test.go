package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRomanToInt(t *testing.T) {
	input := "III"
	expected := 3
	result := romanToInt(input)
	assert.Equal(t, expected, result, "Roman to integer should be equal to expected")

	input = "IV"
	expected = 4
	result = romanToInt(input)
	assert.Equal(t, expected, result, "Roman to integer should be equal to expected")

	input = "IX"
	expected = 9
	result = romanToInt(input)
	assert.Equal(t, expected, result, "Roman to integer should be equal to expected")

	input = "LVIII"
	expected = 58
	result = romanToInt(input)
	assert.Equal(t, expected, result, "Roman to integer should be equal to expected")

	input = "MCMXCIV"
	expected = 1994
	result = romanToInt(input)
	assert.Equal(t, expected, result, "Roman to integer should be equal to expected")
}
