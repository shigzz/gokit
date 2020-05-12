package tree

//TreeNode 二叉树结构体
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Tree interface {
	Insert(num int)
	Search(num int) bool
	Pirnt()
}

//二叉排序树
type CommonBinaryTree struct {
	Val   int
	Left  *CommonBinaryTree
	Right *CommonBinaryTree
}

func CreateCommonTree(num int) *CommonBinaryTree {
	head := &CommonBinaryTree{num, nil, nil}
	return head
}

func (c *CommonBinaryTree) Insert(num int) {
	insert(c, num)
}

func (c *CommonBinaryTree) Search(num int) bool {
	return search(c, num)
}

func (c *CommonBinaryTree) Print() {
}

func insert(head *CommonBinaryTree, num int) {
	if head == nil {
		head = &CommonBinaryTree{num, nil, nil}
	}

	if num > head.Val {
		insert(head.Right, num)
	}
	insert(head.Left, num)
}

func search(head *CommonBinaryTree, num int) bool {
	if head == nil {
		return false
	}
	if head.Val == num {
		return true
	}
	if head.Val > num {
		return search(head.Right, num)
	} else {
		return search(head.Left, num)
	}
}
