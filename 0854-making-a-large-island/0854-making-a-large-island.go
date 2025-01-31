func largestIsland(grid [][]int) int {
    n := len(grid)
    directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // right, up, left, down
    islandSize := map[int]int{}
    islandColor := 2
    maxSize := 0

    // dfs to mark islands and calculate their sizes
    var dfs func(int, int, int) int
    dfs = func (r, c, color int) int {
        if r < 0 || c < 0 || r >= n || c >= n || grid[r][c] != 1 {
            return 0
        }
        grid[r][c] = color
        size := 1
        for _, d := range directions {
            size += dfs(r+d[0], c+d[1], color)
        }
        return size
    }

    // 1, find all the islands and note their sizes
    for r := 0; r < n; r++ {
        for c := 0; c < n; c++ {
            if grid[r][c] == 1 {
                islandSize[islandColor] = dfs(r, c, islandColor)
                maxSize = max(maxSize, islandSize[islandColor])
                islandColor++
            }
        }
    }

    // 2, find the 0
    for r := 0; r < n; r++ {
        for c := 0; c < n; c++ {
            if grid[r][c] == 0 {
                seenIslands := map[int]bool{}
                newSize := 1
                for _, d := range directions {
                    nr, nc := r+d[0], c+d[1]
                    if nr >= 0 && nc >= 0 && nr < n && nc < n && grid[nr][nc] > 1 {
                        color := grid[nr][nc]
                        if !seenIslands[color] {
                            newSize += islandSize[color]
                            seenIslands[color] = true
                        }
                    }
                }
                maxSize = max(maxSize, newSize)
            }
        }
    }

    return maxSize
}

// UTIL
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

// dfs
// time: O(n^2)
// space: O(n^2)