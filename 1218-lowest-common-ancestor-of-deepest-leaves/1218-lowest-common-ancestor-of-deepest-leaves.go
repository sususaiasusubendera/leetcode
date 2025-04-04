/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
    _, lca := helper(root, 0)
    return lca
}

func helper(node *TreeNode, depth int) (int, *TreeNode) {
    if node == nil {
        return depth, nil
    }
    
    leftDepth, leftLCA := helper(node.Left, depth+1)
    rightDepth, rightLCA := helper(node.Right, depth+1)
    
    if leftDepth > rightDepth {
        return leftDepth, leftLCA
    }
    if rightDepth > leftDepth {
        return rightDepth, rightLCA
    }
   
    return leftDepth, node
}

// notice me senpai