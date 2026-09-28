func maxDepth(s string) int {
    md := 0 // max depth
    d := 0 // depth
    for _, c := range s {
        if c == '(' {
            d++
            md = max(md, d)
        } else if c == ')' {
            d--
        }
    }

    return md
}

// string
// time: O(n)
// space: O(1)