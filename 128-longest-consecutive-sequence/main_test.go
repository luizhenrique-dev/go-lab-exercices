package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestConsecutive(t *testing.T) {
	nums := []int{100, 4, 200, 1, 3, 2}
	expected := 4
	result := longestConsecutive(nums)
	assert.Equal(t, expected, result)

	nums = []int{100, 4, 200, 1, 2, -5, 3, -1, -2, -3, -4, 0}
	expected = 10
	result = longestConsecutive(nums)
	assert.Equal(t, expected, result)

	nums = []int{0}
	expected = 1
	result = longestConsecutive(nums)
	assert.Equal(t, expected, result)

	nums = []int{0, 0}
	expected = 1
	result = longestConsecutive(nums)
	assert.Equal(t, expected, result)

	nums = []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	expected = 9
	result = longestConsecutive(nums)
	assert.Equal(t, expected, result)

	nums = []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}
	expected = 19
	result = longestConsecutive(nums)
}

func TestLongestConsecutive2(t *testing.T) {
	nums := []int{100, 4, 200, 1, 3, 2}
	expected := 4
	result := longestConsecutive2(nums)
	assert.Equal(t, expected, result)

	nums = []int{100, 4, 200, 1, 2, -5, 3, -1, -2, -3, -4, 0}
	expected = 10
	result = longestConsecutive2(nums)
	assert.Equal(t, expected, result)

	nums = []int{0}
	expected = 1
	result = longestConsecutive2(nums)
	assert.Equal(t, expected, result)

	nums = []int{0, 0}
	expected = 1
	result = longestConsecutive2(nums)
	assert.Equal(t, expected, result)

	nums = []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	expected = 9
	result = longestConsecutive2(nums)
	assert.Equal(t, expected, result)

	nums = []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}
	expected = 19
	result = longestConsecutive2(nums)
}
