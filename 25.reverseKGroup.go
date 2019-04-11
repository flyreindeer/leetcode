//25.k个一组翻转链表
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	var result *ListNode
	stack := make([]*ListNode, k+1)

	var index int
	cursor := head
	for index = 1; index < k+1 && cursor != nil; index++ {
		stack[index] = cursor
		cursor = cursor.Next
	}

	if index != k+1 {
		return head
	}
	stack[0] = cursor

	result = stack[k]
	var last *ListNode
	for index == k+1 {
		for index = k; index > 0; index-- {
			stack[index].Next = stack[index-1]
		}

		if last != nil {
			last.Next = stack[k]
		}

		last = stack[1]

		for index = 1; index < k+1 && cursor != nil; index++ {
			stack[index] = cursor
			cursor = cursor.Next
		}
		stack[0] = cursor
	}
	return result    
}
