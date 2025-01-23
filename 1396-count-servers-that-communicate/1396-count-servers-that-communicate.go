func countServers(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    
    // count the number of servers in each row and col
    rowServerCount := make([]int, m)
    colServerCount := make([]int, n)
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                rowServerCount[i]++
                colServerCount[j]++
            }
        }
    }

    // count the servers that can communicate
    count := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 && (rowServerCount[i] > 1 || colServerCount[j] > 1) {
                count++
            }
        }
    }

    return count
}

// time: O(mn)
// space: O(m+n)