func minimumDeletions(s string) int {
    n := len(s)
    stack := []byte{}
    ans := 0
    for i := 0; i < n; i++ {
        if len(stack) == 0 { stack = append(stack, s[i]); continue }

        if s[i] == 'a' && stack[len(stack)-1] == 'b' {
            ans++
            stack = stack[:len(stack)-1]
        } else {
            stack = append(stack, s[i])
        }
    }

    return ans
}

// stack, string
// time: O(n)
// space: O(n)