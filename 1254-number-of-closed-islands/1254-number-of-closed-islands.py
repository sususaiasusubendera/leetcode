class Solution:
    def closedIsland(self, grid: list[list[int]]) -> int:
        m, n = len(grid), len(grid[0])

        # mark all the open islands
        # top
        for c in range(n):
            if grid[0][c] == 0:
                bfs(0, c, grid)
        
        # right
        for r in range(m):
            if grid[r][n - 1] == 0:
                bfs(r, n - 1, grid)
        
        # bot
        for c in range(n - 1, -1, -1):
            if grid[m - 1][c] == 0:
                bfs(m - 1, c, grid)

        # left
        for r in range(m - 1, -1, -1):
            if grid[r][0] == 0:
                bfs(r, 0, grid)
        
        # count all the closed islands
        count = 0
        for i in range(m):
            for j in range(n):
                if grid[i][j] == 0:
                    count += 1
                    bfs(i, j, grid)
        
        return count

def bfs(x: int, y: int, grid: list[list[int]]):
    m, n = len(grid), len(grid[0])
    grid[x][y] = 1 # mark
    q = deque([(x, y)]) # pop
    dirs = [(0, -1), (1, 0), (0, 1), (-1, 0)] # up, right, down, left
    while q:
        curr = q.popleft()
        for d in dirs:
            nx, ny = curr[0] + d[0], curr[1] + d[1]
            if 0 <= nx < m and 0 <= ny < n and grid[nx][ny] == 0:
                grid[nx][ny] = 1
                q.append((nx, ny))

# array, bfs
# time: O(mn)
# spacr: O(mn)