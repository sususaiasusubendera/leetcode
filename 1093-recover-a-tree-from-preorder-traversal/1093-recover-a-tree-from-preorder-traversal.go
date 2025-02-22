/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverFromPreorder(traversal string) *TreeNode {
    tree, _ := helper(traversal, 0, 0)
    return tree
}

func helper(s string, index int, depth int) (*TreeNode, int) {
    if index > len(s) {
        return nil, index
    }

    dashCount := 0
    for index < len(s) && s[index] == '-' {
        dashCount++
        index++
    }

    if dashCount != depth {
        return nil, index - dashCount
    }

    start := index
    for index < len(s) && s[index] != '-' {
        index++
    }
    val, _ := strconv.Atoi(s[start:index])

    node := &TreeNode{Val: val}

    node.Left, index = helper(s, index, depth+1)
    node.Right, index = helper(s, index, depth+1)

    return node, index
}

// notice me senpai