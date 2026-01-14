# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def subtreeWithAllDeepest(self, root: Optional[TreeNode]) -> Optional[TreeNode]:
        depth = {None: -1}
        max_depth = 0
        def dfs(node, d, parent = None):
            nonlocal max_depth
            if node:
                max_depth = max(max_depth, d)
                depth[node] = depth[parent] + 1
                dfs(node.left, d + 1,  node)
                dfs(node.right, d + 1, node)
        
        dfs(root, 0)

        def solve(node):
            if not node or depth.get(node, None) == max_depth:
                return node
            left, right = solve(node.left), solve(node.right)
            if left and right:
                return node
            elif left:
                return left
            elif right:
                return right
            else:
                return None
        
        return solve(root)

# binary tree, dfs, hash map, tree
# time: O(n)
# space: O(n)