func findThePrefixCommonArray(A []int, B []int) []int {
    n := len(A)
    f := make(map[int]int) // f: frequency
    result := make([]int, n)
    count := 0
    for i := 0; i < n; i++ {
        f[A[i]]++
        if f[A[i]] == 2 {
            count++
        }

        f[B[i]]++
        if f[B[i]] == 2 {
            count++
        }

        result[i] = count
    }

    return result
}

// frequency map approach