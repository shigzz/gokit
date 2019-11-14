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
