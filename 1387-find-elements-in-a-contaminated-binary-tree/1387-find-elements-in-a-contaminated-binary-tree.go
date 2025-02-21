/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


type FindElements struct {
    seen map[int]bool
}


func Constructor(root *TreeNode) FindElements {
    fe := FindElements{
        seen: map[int]bool{},
    }

    if root != nil {
        fe.recover(root, 0)
    }

    return fe
}


func (this *FindElements) recover(node *TreeNode, val int) {
    if node == nil {
        return
    }

    node.Val = val
    this.seen[val] = true

    this.recover(node.Left, 2*val + 1)
    this.recover(node.Right, 2*val + 2)
}


func (this *FindElements) Find(target int) bool {
    return this.seen[target]
}


/**
 * Your FindElements object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Find(target);
 */

 // time: O(n) (Constructor), O(1) (Find)
 // space: O(n)