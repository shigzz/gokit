package tree

import "testing"

func TestTree(t *testing.T) {
	head := CreateTree([]int{1, 2, 3, 4, 5})
	t.Log(head)
	PrintTree(head)
	//head.Left = &TreeNode{1, nil, nil}
	//MorrisInOrder(head)
}
