/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
    l := leftHeight(root)
    r := rightHeight(root)
    if l == r {
        return pow(2, l) - 1
    } else {
        return 1 + countNodes(root.Left) + countNodes(root.Right)
    }
}

func leftHeight(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return 1 + leftHeight(root.Left)
}

func rightHeight(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return 1 + rightHeight(root.Right)
}

func pow(n, p int) int {
    x := 1
    for i := 0; i < p; i++ {
        x *= n
    }
    return x
}

// binary search - DFS
// time complexity: O(log^2(n))