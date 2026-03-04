# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def sumRootToLeaf(self, r: Optional[TreeNode]) -> int:
        ans = 0
        bin = []

        def dfs(node):
            nonlocal ans

            if node == None: return

            bin.append(node.val)

            if node.left == None and node.right == None:
                ans += binToDec(bin)
                bin.pop() # backtrack
                return

            dfs(node.left)
            dfs(node.right)

            bin.pop() # backtrack
        
        dfs(r)

        return(ans)

# binary tree, dfs, tree
# time: O(n^2)
# space: O(n)

def binToDec(b):
    res = 0
    mul = 1
    for i in range(len(b) - 1, -1, -1):
        if b[i] == 1: res += mul
        mul *= 2
    return res
        