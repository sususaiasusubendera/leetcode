class Solution:
    def createGrid(self, m: int, n: int) -> list[str]:
        ans = []
        for i in range(m):
            if i == m - 1:
                ans.append('.' * n)
            else:
                ans.append('.' + ('#' * (n - 1)))
        return ans

# array, string
# time: O(mn)
# space: O(mn)