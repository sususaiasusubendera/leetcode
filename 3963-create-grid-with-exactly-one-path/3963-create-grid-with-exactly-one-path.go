func createGrid(m int, n int) []string {
    ans := []string{}
    for i := 0; i < m; i++ {
        if i == m-1 {
            ans = append(ans, strings.Repeat(".", n))
        } else {
            ans = append(ans, "." + strings.Repeat("#", n-1))
        }
    }
    return ans
}

// array, string
// time: O(mn)
// space: O(mn)