func maxDepthAfterSplit(seq string) []int {
    ans := []int{}
    depth := 0
    for _, c := range seq {
        if c == '(' {
            depth++
            ans = append(ans, depth%2)
        } else if c == ')' {
            ans = append(ans, depth%2)
            depth--
        }
    }
    return ans
}

// string
// time: O(n)
// space: O(n)