func generateParenthesis(n int) []string {
    temp := make([]byte, 0, 2*n)
    res := []string{}

    var dfs func(open, close int)
    dfs = func(open, close int) {
        if open == 0 && close == 0 {
            res = append(res, string(temp))
            return
        }

        if open > 0 {
            temp = append(temp, '(')
            dfs(open-1, close)
            temp = temp[:len(temp)-1] // backtrack
        }

        if close > open {
            temp = append(temp, ')')
            dfs(open, close-1)
            temp = temp[:len(temp)-1] // backtrack
        }
    }

    dfs(n, n)

    return res
}

// backtracking, string
// time: O(nC_n)
// space: O(nC_n)