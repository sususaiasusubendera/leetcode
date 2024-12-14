/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
    // lDummy: dummy node for nodes that less than x
    // gDummy: dummy node for nodes that greater than or equal to x  
    // l: lDummy pointer
    // g: gDummy pointer
    lDummy, gDummy := &ListNode{}, &ListNode{}
    l, g := lDummy, gDummy
    p := head
    for p != nil {
        if p.Val >= x {
            g.Next = p
            g = g.Next
        } else {
            l.Next = p
            l = l.Next
        }
        p = p.Next
    }
    g.Next = nil
    l.Next = gDummy.Next
    return lDummy.Next
} 