/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    var dfs func(node *TreeNode) int
    dfs = func(node *TreeNode) int {
        if node == nil {
            return 0
        }

        leftHeight := dfs(node.Left)
        if leftHeight == -1 { return -1 }

        rightHeight := dfs(node.Right)
        if rightHeight == -1 { return -1 }

        if abs(leftHeight - rightHeight) > 1 {
            return -1
        }

        return 1 + max(leftHeight, rightHeight)
    }

    if dfs(root) == -1 {
        return false
    }
    return true
}

func abs(n int) int {
    if n < 0 {
        return -n
    }
    return n
}

// binary tree, dfs, tree
// time: O(n)
// space: O(n)