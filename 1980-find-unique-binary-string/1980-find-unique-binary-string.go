func findDifferentBinaryString(nums []string) string {
    n := len(nums)
    m := map[string]bool{}
    for _, num := range nums {
        m[num] = true
    }

    var dfs func(s string) string
    dfs = func(s string) string {
        if len(s) == n {
            if !m[s] {
                return s
            }
            return ""
        }

        addZero := dfs(s + "0")
        if addZero != "" {
            return addZero
        }

        addOne := dfs(s + "1")
        return addOne
    }

    return dfs("")
}

// array, backtracking, hash map, string
// time: O(2^n)
// space: O(n)