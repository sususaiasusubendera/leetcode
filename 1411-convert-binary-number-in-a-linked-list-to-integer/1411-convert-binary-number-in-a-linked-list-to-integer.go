/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
    result := 0
    for head != nil {
        result = result * 2 + head.Val
        head = head.Next
    }

    return result
}

// time: O(n)
// space: O(1)