func findJudge(n int, trust [][]int) int {
    in, out := make([]int, n+1), make([]int, n+1)
    for _, t := range trust {
        in[t[1]]++
        out[t[0]]++
    }

    ans := -1
    for i := 0; i <= n; i++ {
        if in[i] == n-1 && out[i] == 0 {
            ans = i
        }
    }

    return ans
}

// array, graph, hash table
// time: O(n)
// space: O(n)