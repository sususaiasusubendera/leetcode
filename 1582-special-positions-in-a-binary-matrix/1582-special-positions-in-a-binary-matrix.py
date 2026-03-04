class Solution:
    def numSpecial(self, mat: List[List[int]]) -> int:
        m, n = len(mat), len(mat[0])
        ar, ac = [0] * m, [0] * n
        for i in range(0, m):
            for j in range(0, n):
                if mat[i][j] == 1:
                    ar[i] += 1
                    ac[j] += 1
        
        ans = 0
        for i in range(0, m):
            for j in range(0, n):
                if mat[i][j] == 1 and ar[i] == 1 and ac[j] == 1: ans += 1

        return ans

# array, matrix
# time: O(mn)
# space: O(m + n)