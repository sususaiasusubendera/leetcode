func findThePrefixCommonArray(A []int, B []int) []int {
    n := len(A)
    ans := make([]int, n)
    f := map[int]int{}
    total := 0
    for i := 0; i < n; i++ {
        if A[i] == B[i] {
            total++
        } else {
            f[A[i]]++
            f[B[i]]++

            if f[A[i]] == 2 {
                total++
            }

            if f[B[i]] == 2 {
                total++
            }
        }

        ans[i] = total
    }

    return ans
}

// hash map
// time: O(n)
// space: O(n)

// #1 2025/01/14
// #2 2026/05/20