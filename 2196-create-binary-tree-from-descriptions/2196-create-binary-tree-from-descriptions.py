# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def createBinaryTree(self, descriptions: List[List[int]]) -> Optional[TreeNode]:
        nodes = {}
        is_child = {}
        for parent, child, is_left in descriptions:
            if parent not in nodes:
                nodes[parent] = TreeNode(parent)
            
            if child not in nodes:
                nodes[child] = TreeNode(child)
            
            if is_left:
                nodes[parent].left = nodes[child]
            else:
                nodes[parent].right = nodes[child]
            
            is_child[child] = True
        
        for node in nodes:
            if node not in is_child:
                return nodes[node]
        
        return None

# binary tree, hash map, tree
# time: O(n)
# space: O(n)