package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseOnlyLettersCase1(t *testing.T) {
	s := "ab-cd"
	expected := "dc-ba"
	result := reverseOnlyLetters(s)
	assert.Equal(t, expected, result, "Reverse only letters should be equal to expected")
}

func TestReverseOnlyLettersCase2(t *testing.T) {
	s := "a-bC-dEf-ghIj"
	expected := "j-Ih-gfE-dCba"
	result := reverseOnlyLetters(s)
	assert.Equal(t, expected, result, "Reverse only letters should be equal to expected")
}

func TestReverseOnlyLettersCase3(t *testing.T) {
	s := "Test1ng-Leet=code-Q!"
	expected := "Qedo1ct-eeLg=ntse-T!"
	result := reverseOnlyLetters(s)
	assert.Equal(t, expected, result, "Reverse only letters should be equal to expected")
}

func TestReverseOnlyLettersAlgorithm2Case1(t *testing.T) {
	s := "ab-cd"
	expected := "dc-ba"
	result := reverseOnlyLetters2(s)
	assert.Equal(t, expected, result, "Reverse only letters should be equal to expected")
}

func TestReverseOnlyLettersAlgorithm2Case2(t *testing.T) {
	s := "a-bC-dEf-ghIj"
	expected := "j-Ih-gfE-dCba"
	result := reverseOnlyLetters2(s)
	assert.Equal(t, expected, result, "Reverse only letters should be equal to expected")
}

func TestReverseOnlyLettersAlgorithm2Case3(t *testing.T) {
	s := "Test1ng-Leet=code-Q!"
	expected := "Qedo1ct-eeLg=ntse-T!"
	result := reverseOnlyLetters2(s)
	assert.Equal(t, expected, result, "Reverse only letters should be equal to expected")
}
