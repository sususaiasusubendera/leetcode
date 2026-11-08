func countServers(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    rowServers := make([]int, m)
    colServers := make([]int, n)
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                rowServers[i]++
                colServers[j]++
            }
        }
    }

    count := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 && (rowServers[i] > 1 || colServers[j] > 1) {
                count++
            }
        }
    }

    return count
}

// array, counting, matrix
// time: O(mn)
// space: O(m + n)