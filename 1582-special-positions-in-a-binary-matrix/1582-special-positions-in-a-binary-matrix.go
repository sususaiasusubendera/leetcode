func numSpecial(mat [][]int) int {
    m, n := len(mat), len(mat[0])
    ar, ac := make([]int, m), make([]int, n) // freq row and col
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if mat[i][j] == 1 {
                ar[i]++
                ac[j]++
            }
        }
    }

    ans := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if mat[i][j] == 1 && ar[i] == 1 && ac[j] == 1 { ans++ }
        }
    }

    return ans
}

// array, matrix
// time: O(mn)
// space: O(m + n)