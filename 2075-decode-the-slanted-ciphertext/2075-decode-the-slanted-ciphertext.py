class Solution:
    def decodeCiphertext(self, encodedText: str, rows: int) -> str:
        m, n = rows, len(encodedText) // rows

        grid = [[' '] * n for _ in range(m)]

        r, c = 0, 0
        for i in range(len(encodedText)):
            grid[r][c] = encodedText[i]
            c += 1
            if c == n:
                c = 0
                r += 1
        
        ans = []
        for j in range(n):
            r, c = 0, j
            while r < m and c < n:
                ans.append(grid[r][c])
                r += 1
                c += 1
        
        idx = len(ans) - 1
        while idx >= 0 and ans[idx] == ' ':
            idx -= 1
        
        return ''.join(ans[:idx+1])

# string
# time: O(n)
# space: O(n)