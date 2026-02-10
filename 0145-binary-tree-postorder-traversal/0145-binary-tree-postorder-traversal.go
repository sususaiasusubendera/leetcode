/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
    ans := []int{}

    var dfs func(node *TreeNode)
    dfs = func(node *TreeNode) {
        if node == nil { return }

        dfs(node.Left)
        dfs(node.Right)
        ans = append(ans, node.Val)
    }

    dfs(root)

    return ans
}

// binary tree, dfs, stack, tree
// time: O(n)
// space: O(n)