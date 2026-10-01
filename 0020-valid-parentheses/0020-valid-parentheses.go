func isValid(s string) bool {
    pairs := map[rune]rune{
        '(': ')',
        '{': '}',
        '[': ']',
    }

    if len(s) % 2 == 1 {
        return false
    }

    stack := []rune{}
    for _, c := range s {
        if _, ok := pairs[c]; ok { // c is open bracket
            stack = append(stack, c)
        } else if len(stack) == 0 || c != pairs[stack[len(stack)-1]] {
            // close bracket appears before any open bracket or top != current close bracket
            return false
        } else {
            stack = stack[:len(stack)-1] // pop
        }
    }
    
    return len(stack) == 0 // all open brackets have their pairs
}

// stack, string
// time: O(n)
// space: O(m)