/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func modifiedList(nums []int, head *ListNode) *ListNode {
    m := map[int]struct{}{}
    for i := 0; i < len(nums); i++ {
        m[nums[i]] = struct{}{}
    }

    dummy := &ListNode{Next: head}
    prev := dummy
    curr := head
    for curr != nil {
        if _, exists := m[curr.Val]; exists {
            prev.Next = curr.Next
        } else {
            prev = curr
        }
        curr = curr.Next
    }

    return dummy.Next 
}

// array, hash map, linked list
// time: O(n + m)
// space: O(n)