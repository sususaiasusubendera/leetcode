func isValid(s string) bool {
    if len(s) == 0 || len(s) % 2 == 1 {
        return false
    }

    pairs := map[byte]byte{
        '{': '}',
        '(': ')',
        '[': ']',
    }
    stack := []byte{}

    for i := 0; i < len(s); i++ {
        if _, exist := pairs[s[i]]; exist {
            stack = append(stack, s[i])
        } else if len(stack) == 0 || s[i] != pairs[stack[len(stack)-1]] {
            return false
        } else {
            stack = stack[:len(stack)-1]
        }
    }
    return len(stack) == 0
}