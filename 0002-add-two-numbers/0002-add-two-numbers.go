/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    remainder := 0
    dummy := &ListNode{}
    current := dummy

    for l1 != nil || l2 != nil || remainder != 0 {
        val1 := 0
        if l1 != nil {
            val1 = l1.Val
            l1 = l1.Next
        }

        val2 := 0
        if l2 != nil {
            val2 = l2.Val
            l2 = l2.Next
        }

        sum := val1 + val2 + remainder
        remainder = sum / 10
        current.Next = &ListNode{
            Val: sum % 10,
        }
        current = current.Next
    }
    return dummy.Next
}