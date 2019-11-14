package linklist

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{
		nums[0],
		nil,
	}
	head.Next = CreateList(nums[1:])
	return head
}

func CreateCycleList(nums []int, pos int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	if pos < 0 {
		return CreateList(nums)
	}
	head := &ListNode{nums[0], nil}
	pre := head
	var loc *ListNode
	if pos == 0 {
		loc = head
	}
	for i := 1; i < len(nums); i++ {
		pre.Next = &ListNode{nums[i], nil}
		pre = pre.Next
		if i == pos {
			loc = pre
		}
	}
	pre.Next = loc
	return head
}
