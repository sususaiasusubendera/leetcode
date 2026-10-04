class Solution:
    def generateParenthesis(self, n: int) -> list[str]:
        temp = []
        res = []

        def dfs(open, close):
            if not open and not close:
                res.append("".join(temp))
                return
            
            if open > 0:
                temp.append("(")
                dfs(open - 1, close)
                temp.pop()
            
            if close > open:
                temp.append(")")
                dfs(open, close - 1)
                temp.pop()
        
        dfs(n, n)
        return res

# backtracking, string
# time: O(nC_n)
# space: O(nC_n)