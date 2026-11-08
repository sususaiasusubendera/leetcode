class Solution:
    def countServers(self, grid: list[list[int]]) -> int:
        m, n = len(grid), len(grid[0])
        row_servers = [0] * m
        col_servers = [0] * n
        for i in range(m):
            for j in range(n):
                if grid[i][j]:
                    row_servers[i] += 1
                    col_servers[j] += 1
        
        count = 0
        for i in range(m):
            for j in range(n):
                if grid[i][j] and (row_servers[i] > 1 or col_servers[j] > 1):
                    count += 1
        
        return count

# array, counting, matrix
# time: O(mn)
# space: O(m + n)