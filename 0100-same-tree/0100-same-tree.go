/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil {
        return true
    }
    if (p == nil && q != nil) || (p != nil && q == nil) {
        return false
    }
    if p.Val != q.Val {
        return false
    }

    leftCheck := isSameTree(p.Left, q.Left)
    rightCheck := isSameTree(p.Right, q.Right)

    return leftCheck && rightCheck
}