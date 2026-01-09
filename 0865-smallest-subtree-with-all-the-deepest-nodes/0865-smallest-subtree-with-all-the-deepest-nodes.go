/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
    maxDepth := 0
    depth := map[*TreeNode]int{}

    // dfs find max depth
    var dfs func(node *TreeNode, d int)
    dfs = func(node *TreeNode, d int) {
        if node == nil {
            return
        }

        maxDepth = max(maxDepth, d)
        depth[node] = d

        dfs(node.Left, d+1)
        dfs(node.Right, d+1)
    }

    dfs(root, 1)

    var solve func(node *TreeNode) *TreeNode
    solve = func(node *TreeNode) *TreeNode {
        if node == nil || depth[node] == maxDepth {
            return node
        }

        left, right := solve(node.Left), solve(node.Right) 
        if left != nil && right != nil {
            return node
        } else if left != nil {
            return left
        } else if right != nil {
            return right
        }
        return nil
    }

    return solve(root)
}

// binary tree, dfs, hash map, tree
// time: O(n)
// space: O(n)