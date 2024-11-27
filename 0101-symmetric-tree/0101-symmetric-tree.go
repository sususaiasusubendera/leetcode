/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }

    // using anonym function
    var isMirror func (t1, t2 *TreeNode) bool
    isMirror = func (t1, t2 *TreeNode) bool {
        if t1 == nil && t2 == nil {
            return true
        }
        if t1 == nil || t2 == nil {
            return false
        }
        if t1.Val != t2.Val {
            return false
        }
        return isMirror(t1.Left, t2.Right) && isMirror(t1.Right, t2.Left)
    }

    return isMirror(root.Left, root.Right)
}