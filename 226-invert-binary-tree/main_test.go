package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvertTreeCase1(t *testing.T) {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{
			Val:   7,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 9},
		},
	}
	expected := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   7,
			Left:  &TreeNode{Val: 9},
			Right: &TreeNode{Val: 6},
		},
		Right: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 1},
		},
	}
	result := invertTree(root)
	assert.Equal(t, expected, result, "Inverted tree should be equal to expected")
}

func TestInvertTreeCase2(t *testing.T) {
	root := &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 3},
	}
	expected := &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 3},
		Right: &TreeNode{Val: 1},
	}
	result := invertTree(root)
	assert.Equal(t, expected, result, "Inverted tree should be equal to expected")
}
