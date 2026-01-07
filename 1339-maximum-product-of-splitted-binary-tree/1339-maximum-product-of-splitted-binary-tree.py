# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def maxProduct(self, root: Optional[TreeNode]) -> int:
        MOD = 1_000_000_007

        def dfs_sum(root):
            if not root:
                return 0
            return root.val + dfs_sum(root.left) + dfs_sum(root.right)
        
        total = dfs_sum(root)
        max_product = 0

        def solve(root):
            nonlocal max_product
            if not root:
                return 0

            sum = root.val + solve(root.left) + solve(root.right)
            max_product = max(max_product, sum * (total - sum))

            return sum

        solve(root)
        return max_product % MOD

# dfs, binary tree, tree
# time: O(n)
# space: O(h)