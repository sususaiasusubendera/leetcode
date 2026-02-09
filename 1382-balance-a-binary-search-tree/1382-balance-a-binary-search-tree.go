/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func balanceBST(root *TreeNode) *TreeNode {
    // construct inorder array val
    inorder := []int{}
    var dfs func(node *TreeNode)
    dfs = func(node *TreeNode) {
        if node == nil { return }

        dfs(node.Left)
        inorder = append(inorder, node.Val)
        dfs(node.Right)
    }

    dfs(root)

    // build the tree using inorder array
    var build func(l, r int) *TreeNode
    build = func(l, r int) *TreeNode {
        if l > r { return nil }

        mid := l + ((r - l) / 2)
        node := &TreeNode{inorder[mid], build(l, mid-1), build(mid+1, r)}

        return node
    }

    return build(0, len(inorder)-1)
}

// binary tree, binary search tree, dfs, divide and conquer, senior(?), tree
// time: O(n)
// space: O(n)