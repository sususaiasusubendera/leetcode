func findMissingAndRepeatedValues(grid [][]int) []int {
    n := len(grid)
    var repeated, missing int
    seen := map[int]int{}
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if _, exists := seen[grid[i][j]]; !exists {
                seen[grid[i][j]] = 1
                continue
            }

            repeated = grid[i][j]
        }
    }

    n *= n
    for i := 1; i <= n; i++ {
        if _, exists := seen[i]; !exists {
            missing = i
        }
    }

    return []int{repeated, missing}
}

// time: O(n^2)
// space: O(n)