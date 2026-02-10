# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def balanceBST(self, root: Optional[TreeNode]) -> Optional[TreeNode]:
        inorder = []
        def dfs(node):
            if node == None: return

            dfs(node.left)
            inorder.append(node.val)
            dfs(node.right)
        dfs(root)

        def build(l, r):
            if l > r: return None

            mid = l + ((r - l) // 2)
            node = TreeNode(inorder[mid], build(l, mid - 1), build(mid + 1, r))
            return node
        
        return build(0, len(inorder)-1)

# binary tree, binary search tree, dfs, divide and conquer, senior(?), tree
# time: O(n)
# space: O(n)