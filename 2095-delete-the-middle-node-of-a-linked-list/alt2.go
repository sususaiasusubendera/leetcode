/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {
    // n == 1
    if head.Next == nil {
        return nil
    }

    slow := head
    fast := slow.Next.Next
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }

    // delete
    slow.Next = slow.Next.Next

    return head
}

// floyd's tortoise & hare
// time: O(n)
// space: O(1)

// AMAZING BEAUTIFUL solution
