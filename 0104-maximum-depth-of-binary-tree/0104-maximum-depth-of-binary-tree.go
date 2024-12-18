/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }

    leftHeight := maxDepth(root.Left)
    rightHeight := maxDepth(root.Right)
    
    return 1 + max(leftHeight, rightHeight)
}

func max(a, b int) int {
    if a >= b {
        return a
    } else {
        return b
    }
}