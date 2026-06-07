/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func createBinaryTree(descriptions [][]int) *TreeNode {
    nodes := map[int]*TreeNode{} // set nodes (1 on 1)
    isChilds := map[int]bool{}
    for _, d := range descriptions {
        parent, child, isLeft := d[0], d[1], d[2]

        if nodes[parent] == nil {
            nodes[parent] = &TreeNode{Val: parent}
        }

        if nodes[child] == nil {
            nodes[child] = &TreeNode{Val: child}
        }

        if isLeft == 1 {
            nodes[parent].Left = nodes[child]
        } else {
            nodes[parent].Right = nodes[child]
        }

        isChilds[child] = true
    }

    for val, node := range nodes {
        if !isChilds[val] {
            return node
        }
    }

    return nil
}

// binary tree, hash map, tree
// time: O(n)
// space: O(n)

// note: a node is root if it hasn't became a child