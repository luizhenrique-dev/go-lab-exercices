package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindromeNumber(t *testing.T) {
	assert.True(t, isPalindrome(121))
	assert.False(t, isPalindrome(-121))
	assert.False(t, isPalindrome(10))
	assert.True(t, isPalindrome(0))
	assert.True(t, isPalindrome(323))
	assert.False(t, isPalindrome(54213))
	assert.True(t, isPalindrome(9876789))
	assert.False(t, isPalindrome(-9876789))
	assert.False(t, isPalindrome(-1))
	assert.False(t, isPalindrome(98767891230))
}
