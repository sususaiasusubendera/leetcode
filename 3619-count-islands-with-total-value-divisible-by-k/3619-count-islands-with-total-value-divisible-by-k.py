class Solution:
    def countIslands(self, grid: List[List[int]], k: int) -> int:
        m, n = len(grid), len(grid[0])
        dirs = [(0, 1), (1, 0), (0, -1), (-1, 0)] # right, down, left, up

        def bfs(r: int, c: int) -> int:
            q = deque([(r, c)])
            total = grid[r][c]
            grid[r][c] = 0 # mark visited
            while q:
                cr, cc = q.popleft()
                for dr, dc in dirs:
                    nr, nc = cr + dr, cc + dc
                    if 0 <= nr < m and 0 <= nc < n and grid[nr][nc] > 0:
                        total += grid[nr][nc]
                        grid[nr][nc] = 0 # mark visited
                        q.append((nr, nc))
            return total

        ans = 0
        for i in range(m):
            for j in range(n):
                if grid[i][j] and bfs(i, j) % k == 0:
                    ans += 1
        
        return ans

# array, bfs, matrix
# time: O(mn)
# space: O(mn)