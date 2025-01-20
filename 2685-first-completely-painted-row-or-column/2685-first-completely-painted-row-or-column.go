func firstCompleteIndex(arr []int, mat [][]int) int {
    m, n := len(mat), len(mat[0])
    pos := make([][2]int, (m * n) + 1) // pos: position

    // pre-process the positions of the values in the matrix
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            pos[mat[i][j]][0] = i
            pos[mat[i][j]][1] = j
        }
    }

    // count rowFreq and colFreq based on pre-processed pos 
    rowFreq := make([]int, m)
    colFreq := make([]int, n)
    for i := 0; i < len(arr); i++ {
        rowFreq[pos[arr[i]][0]]++
        colFreq[pos[arr[i]][1]]++

        // if the rowFreq == number of col, or vice versa return the current index
        if rowFreq[pos[arr[i]][0]] == n || colFreq[pos[arr[i]][1]] == m {
            return i
        }
    }

    return -1
}

// time: O(mn + k)
// space: O(mn)