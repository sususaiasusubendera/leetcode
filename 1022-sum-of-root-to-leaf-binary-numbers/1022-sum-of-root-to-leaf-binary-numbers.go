/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumRootToLeaf(root *TreeNode) int {
	ans := 0
	bin := []int{}

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		bin = append(bin, node.Val)

		if node.Left == nil && node.Right == nil {
			ans += binToDec(bin)
            bin = bin[:len(bin)-1]
            return
		}

        dfs(node.Left)
        dfs(node.Right)

        bin = bin[:len(bin)-1]
	}

    dfs(root)
	return ans
}

// binary tree, dfs, tree
// time: O(n^2)
// space: O(n)

func binToDec(b []int) int {
	res := 0
	mul := 1
	for i := len(b) - 1; i >= 0; i-- {
		if b[i] == 1 {
			res += mul
		}
		mul *= 2
	}
	return res
}