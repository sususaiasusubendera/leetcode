func minAddToMakeValid(s string) int {
    stack := []rune{}
    for _, c := range s {
        if len(stack) > 0 && stack[len(stack)-1] == '(' {
            if c == ')' {
                stack = stack[:len(stack)-1] // pop
                continue
            }
        }
        stack = append(stack, c)
    }

    return len(stack)
}

// greedy, stack, string
// time: O(n)
// space: O(n)