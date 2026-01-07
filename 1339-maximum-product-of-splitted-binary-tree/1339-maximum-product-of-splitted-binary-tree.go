/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxProduct(root *TreeNode) int {
    const MOD = 1_000_000_007

    var dfsSum func(root *TreeNode) int
    dfsSum = func(root *TreeNode) int {
        if root == nil {
            return 0
        }

        return root.Val + dfsSum(root.Left) + dfsSum(root.Right)
    }

    total := dfsSum(root)
    maxProduct := 0

    var solve func(root *TreeNode) int
    solve = func(root *TreeNode) int {
        if root == nil {
            return 0
        }

        sum := root.Val + solve(root.Left) + solve(root.Right)
        product := sum * (total - sum)
        maxProduct = max(maxProduct, product)

        return sum
    }

    solve(root)
    return maxProduct % MOD
}

// binary tree, dfs, tree
// time: O(n)
// space: O(h) (recursion stack -> tree's height)
