func scoreOfParentheses(s string) int {
    var solve func(left, right int) int
    solve = func(left, right int) int {
        if left+1 == right { // base case: ()
            return 1
        }

        depth := 0
        for i := left; i <= right; i++ {
            if s[i] == '(' {
                depth++
            } else { // ')'
                depth--
            }

            if depth == 0 {
                if i == right { // (A)
                    return 2 * solve(left+1, right-1)
                } else { // AB
                    return solve(left, i) + solve(i+1, right)
                }
            }
        }

        return 0
    }

    return solve(0, len(s)-1)
}

// divide and conquer, recursive, string
// time: O(n)
// space: O(n)