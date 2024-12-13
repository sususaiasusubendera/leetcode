/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    
    dummy := &ListNode{0, head}
    p, q := head, dummy
    for p != nil && p.Next != nil {
        if p.Val == p.Next.Val {
            for p.Next != nil && p.Val == p.Next.Val {
                p = p.Next
            }
            q.Next = p.Next
            p = p.Next
        } else {
            p = p.Next
            q = q.Next
        }
    }
    return dummy.Next
}