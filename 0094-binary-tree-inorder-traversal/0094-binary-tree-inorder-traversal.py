# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def inorderTraversal(self, root: Optional[TreeNode]) -> List[int]:
        ans = []

        def dfs(node):
            if node == None: return

            dfs(node.left)
            ans.append(node.val)
            dfs(node.right)

        dfs(root)

        return ans

# binary tree, dfs, stack, tree
# time: O(n)
# space: O(n)