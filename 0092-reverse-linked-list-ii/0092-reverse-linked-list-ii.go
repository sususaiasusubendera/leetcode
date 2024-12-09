/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
    if head == nil || left == right {
        return head
    }

    dummy := &ListNode{0, head}
    prev := dummy
    for i := 0; i < left - 1; i++ {
        prev = prev.Next
    }

    curr := prev.Next
    var prevNode *ListNode
    for i := 0; i <= right - left; i++ {
        nextNode := curr.Next
        curr.Next = prevNode
        prevNode = curr
        curr = nextNode
    }

    prev.Next.Next = curr
    prev.Next = prevNode

    return dummy.Next
}