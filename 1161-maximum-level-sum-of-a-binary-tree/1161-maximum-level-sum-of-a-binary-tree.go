/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxLevelSum(root *TreeNode) int {
    maxSum, ans, level := int(-1 << 31), 0, 0

    queue := []*TreeNode{root}

    for len(queue) > 0 {
        level++
        sumPerLevel := 0
        levelSize := len(queue)
        for i := 0; i < levelSize; i++ {
            node := queue[0]
            queue = queue[1:] // pop

            sumPerLevel += node.Val

            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }

        if sumPerLevel > maxSum {
            maxSum = sumPerLevel
            ans = level
        }
    }

    return ans
}

// bfs, binary tree, tree
// time: O(n)
// space: O(n)