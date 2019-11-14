package tree

import "testing"

func testTree(t *testing.T) {
	head := CreateTree([]int{1, 2, 3, 4, 5})
	PrintTree(head)
}
