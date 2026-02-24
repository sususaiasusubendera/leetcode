class Solution:
    def sumRootToLeaf(self, r: Optional[TreeNode]) -> int:
        return (f:=lambda n,q=0:n and (f(n.left,q:=q*2+n.val)+f(n.right,q) or q) or 0)(r)

# notice me senpai
# solution from solutions (MikPosp)