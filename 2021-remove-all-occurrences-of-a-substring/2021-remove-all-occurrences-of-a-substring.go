func removeOccurrences(s string, part string) string {
    stack := ""
    for i := 0; i < len(s); i++ {
        stack += string(s[i])
        if len(stack) >= len(part) && stack[len(stack)-len(part):] == part {
            stack = stack[:len(stack)-len(part)]
        }
    }

    return stack
}

// stack approach
// time: O(mn)
// space: O(n)