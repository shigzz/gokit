package tree

//Node 二叉树结构体
type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

//CreateTree 层序方式
//接收int数组创建二叉树，-1为nil节点
func CreateTree(nums []int) *Node {
	if len(nums) == 0 || nums[0] == -1 {
		return nil
	}
	nodes := make([]*Node, len(nums))
	for i := range nums {
		if nums[i] != -1 {
			nodes[i] = &Node{nums[i], nil, nil}
			if i != 0 {
				if i%2 == 1 {
					nodes[i/2].Left = nodes[i]
				} else {
					nodes[i/2-1].Right = nodes[i]
				}
			}
		}
	}
	return nodes[0]
}

//CreateTreeByInterface 接收任何对象，”null“为nil节点
func CreateTreeByInterface(nums []interface{}) *Node {
	if len(nums) == 0 || nums[0] == "null" {
		return nil
	}
	nodes := make([]*Node, len(nums))
	for i := range nums {
		if nums[i] != "null" {
			//nodes[i] = &Node{nums[i], nil, nil}
			n, _ := nums[i].(int)
			nodes[i] = &Node{n, nil, nil}
			if i != 0 {
				if i%2 == 1 {
					nodes[i/2].Left = nodes[i]
				} else {
					nodes[i/2-1].Right = nodes[i]
				}
			}
		}
	}
	return nodes[0]
}
