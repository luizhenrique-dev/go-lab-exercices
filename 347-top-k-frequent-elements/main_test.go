package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopKFrequentCase1(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	expected := []int{1, 2}
	result := topKFrequent(nums, k)
	assert.Equal(t, expected, result, "Top k frequent elements should be equal to expected")
}

func TestTopKFrequentCase2(t *testing.T) {
	nums := []int{1}
	k := 1
	expected := []int{1}
	result := topKFrequent(nums, k)
	assert.Equal(t, expected, result, "Top k frequent elements should be equal to expected")
}

func TestTopKFrequentCase3(t *testing.T) {
	nums := []int{1, 1, 1, 1, 2, 2, 3, 3, 3, 3, 3, 2, 5, 5, 5, 6}
	k := 2
	expected := []int{3, 1}
	result := topKFrequent(nums, k)
	assert.Equal(t, expected, result, "Top k frequent elements should be equal to expected")
}

func TestTopKFrequentCase1Algorithm2(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	expected := []int{1, 2}
	result := topKFrequent2(nums, k)
	assert.Equal(t, expected, result, "Top k frequent elements should be equal to expected")
}

func TestTopKFrequentCase2Algorithm2(t *testing.T) {
	nums := []int{1}
	k := 1
	expected := []int{1}
	result := topKFrequent2(nums, k)
	assert.Equal(t, expected, result, "Top k frequent elements should be equal to expected")
}

func TestTopKFrequentCase3Algorithm2(t *testing.T) {
	nums := []int{1, 1, 1, 1, 2, 2, 3, 3, 3, 3, 3, 2, 5, 5, 5, 6}
	k := 2
	expected := []int{3, 1}
	result := topKFrequent2(nums, k)
	assert.Equal(t, expected, result, "Top k frequent elements should be equal to expected")
}
