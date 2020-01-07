package tree

import "fmt"

func preOrderTree(root *TreeNode) {
	if root != nil {
		fmt.Print(root.Val, " ")
		preOrderTree(root.Left)
		preOrderTree(root.Right)
	}
}

func inOrderTree(root *TreeNode) {
	if root != nil {
		inOrderTree(root.Left)
		fmt.Print(root.Val, " ")
		inOrderTree(root.Right)
	}
}

func postOrderTree(root *TreeNode) {
	if root != nil {
		postOrderTree(root.Left)
		postOrderTree(root.Right)
		fmt.Print(root.Val, " ")
	}
}

//MorrisInOrder morris解法，空间复杂度为O(n)
func MorrisInOrder(root *TreeNode) {
	cur := root
	var pre *TreeNode
	for cur != nil {
		if cur.Left != nil {
			pre = cur.Left
			for pre.Right != nil && pre.Right != cur {
				pre = pre.Right
			}
			if pre.Right == nil {
				pre.Right = cur
				cur = cur.Left
			} else {
				pre.Right = nil
				fmt.Println(cur.Val, " ")
				cur = cur.Right
			}
		} else {
			fmt.Println(cur.Val, " ")
			cur = cur.Right
		}
	}
}

//PrintTree 三种方式遍历二叉树
func PrintTree(root *TreeNode) {
	fmt.Print("PreOrder: ")
	preOrderTree(root)
	fmt.Println()
	fmt.Print("InOrder: ")
	inOrderTree(root)
	fmt.Println()
	fmt.Print("PostOrder: ")
	postOrderTree(root)
	fmt.Println()
}
