/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {
    n := 0
    curr := head
    for curr != nil {
        n++
        curr = curr.Next
    }

    if n == 2 {
        head.Next = nil
        return head
    } else if n == 1 {
        return nil
    }

    mid := n / 2
    prev := head
    for i := 0; i < mid - 1; i++ {
        prev = prev.Next
    }

    prev.Next = prev.Next.Next

    return head
}

// linked list
// time: O(n)
// space: O(1)