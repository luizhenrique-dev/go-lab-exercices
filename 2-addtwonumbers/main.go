package main

import (
	"fmt"
	"math/big"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// main function
// Add Two Numbers
// You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit.
// Add the two numbers and return the sum as a linked list.
func main() {
	arrL1 := []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	l1 := arrayToListNode(arrL1)
	arrL2 := []int{5, 6, 4}
	l2 := arrayToListNode(arrL2)
	result := addTwoNumbers(l1, l2)
	printList(result)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1StrCh := make(chan string)
	l2StrCh := make(chan string)

	go listToString(l1, l1StrCh)
	go listToString(l2, l2StrCh)

	l1Str := <-l1StrCh
	l2Str := <-l2StrCh

	l1Num := stringToBigInt(l1Str)
	l2Num := stringToBigInt(l2Str)

	sum := new(big.Int).Add(l1Num, l2Num)
	return stringToListNode(sum.String())
}

func listToString(list *ListNode, result chan<- string) {
	var str string
	for current := list; current != nil; current = current.Next {
		str += strconv.Itoa(current.Val)
	}
	result <- reverseString(str)
	close(result)
}

func stringToBigInt(s string) *big.Int {
	num := new(big.Int)
	num.SetString(s, 10)
	return num
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func stringToListNode(s string) *ListNode {
	var head *ListNode
	for _, char := range s {
		num, _ := strconv.Atoi(string(char))
		head = &ListNode{Val: num, Next: head}
	}
	return head
}

func arrayToListNode(arr []int) *ListNode {
	var prev *ListNode
	for i := len(arr) - 1; i >= 0; i-- {
		node := &ListNode{
			Val:  arr[i],
			Next: prev,
		}
		prev = node
	}
	return prev
}

func printList(node *ListNode) string {
	var result string
	for node != nil {
		result += strconv.Itoa(node.Val) + " -> "
		node = node.Next
	}
	result += "nil"
	fmt.Println(result)
	return result
}
