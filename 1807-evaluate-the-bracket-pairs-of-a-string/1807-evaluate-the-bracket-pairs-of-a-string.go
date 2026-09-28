func evaluate(s string, knowledge [][]string) string {
    m := map[string]string{}
    for _, k := range knowledge {
        m[k[0]] = k[1]
    }

    ans := strings.Builder{}
    start := -1
    for i, c := range s {
        if c == '(' {
            start = i
        } else if c == ')' {
            if val, ok := m[s[start+1:i]]; ok {
                ans.WriteString(val)
            } else {
                ans.WriteRune('?')
            }
            start = -1
        } else if start == -1 {
            ans.WriteRune(c)
        }
    }
    return ans.String()
}

// array, hash map, string
// time: O(n + m)
// space: O(n + m)