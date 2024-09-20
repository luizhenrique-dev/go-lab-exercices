package main

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t *testing.T) {
	arrL1 := []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	l1 := arrayToListNode(arrL1)
	arrL2 := []int{5, 6, 4}
	l2 := arrayToListNode(arrL2)
	expected := []int{6, 6, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	result := addTwoNumbers(l1, l2)
	assert.Equal(t, expected, listNodeToArray(result))
}

func TestListToString(t *testing.T) {
	list := arrayToListNode([]int{1, 2, 3})
	resultCh := make(chan string)
	go listToString(list, resultCh)
	result := <-resultCh
	assert.Equal(t, "321", result)
}

func TestStringToBigInt(t *testing.T) {
	result := stringToBigInt("12345678901234567890")
	expected := new(big.Int)
	expected.SetString("12345678901234567890", 10)
	assert.Equal(t, expected, result)
}

func TestReverseString(t *testing.T) {
	result := reverseString("abcdef")
	assert.Equal(t, "fedcba", result)
}

func TestStringToListNode(t *testing.T) {
	result := stringToListNode("123")
	expected := arrayToListNode([]int{3, 2, 1})
	assert.Equal(t, expected, result)
}

func TestArrayToList(t *testing.T) {
	arr := []int{1, 2, 3}
	result := arrayToListNode(arr)
	expected := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	assert.Equal(t, expected, result)
}

func TestPrintList(t *testing.T) {
	list := arrayToListNode([]int{1, 2, 3})
	expected := "1 -> 2 -> 3 -> nil"
	assert.Equal(t, expected, printList(list))
}

// Helper functions
func listNodeToArray(list *ListNode) []int {
	var result []int
	for list != nil {
		result = append(result, list.Val)
		list = list.Next
	}
	return result
}
