/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    if k == 0 || head == nil || head.Next == nil {
        return head
    }

    tail, length := head, 1
    for tail.Next != nil {
        tail = tail.Next
        length++
    }
    
    k %= length
    if k == 0 {
        return head
    }

    newTailPos := length - k
    newTail := head
    for i := 1; i < newTailPos; i++ {
        newTail = newTail.Next
    }

    newHead := newTail.Next
    newTail.Next = nil
    tail.Next = head
    
    return newHead
}