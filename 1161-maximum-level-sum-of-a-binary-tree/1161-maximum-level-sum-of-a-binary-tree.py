# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def maxLevelSum(self, root: Optional[TreeNode]) -> int:
        max_sum, ans, level = -100000, 0, 0
        q = deque([root])
        while q:
            level += 1
            sum_per_level = 0
            level_size = len(q)
            for _ in range(level_size):
                curr = q.popleft()
                sum_per_level += curr.val
                
                if curr.left:
                    q.append(curr.left)
                if curr.right:
                    q.append(curr.right)

            if sum_per_level > max_sum:
                max_sum = sum_per_level
                ans = level
        
        return ans

# bfs, binary tree, tree
# time: O(n)
# space: O(n)