func largestOverlap(img1 [][]int, img2 [][]int) int {
    ans := 0
    n := len(img1)
    for i := -(n - 1); i <= n - 1; i++ {
        for j := -(n - 1); j <= n - 1; j++ {
            ans = max(ans, countOne(i, j, img1, img2))
        }
    }

    return ans
}

func countOne(sr, sc int, img1, img2 [][]int) int {
    overlap := 0
    n := len(img1)
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            ni := i + sr
            nj := j + sc
            if 0 <= ni && ni < n && 0 <= nj && nj < n {
                if img1[ni][nj] == 1 && img2[i][j] == 1 {
                    overlap++
                }
            } 
        }
    }

    return overlap
}

// array, brute force, matrix
// time: O(n^4)
// space: O(1)