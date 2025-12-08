func countTriples(n int) int {
    count := 0
    for a := 1; a <= n; a++ {
        for b := 1; b <= n; b++ {
            c := int(math.Sqrt(float64(a*a + b*b)))
            if c <= n && c*c == a*a + b*b { count++ }
        }
    }

    return count
}

// math
// time: O(n^2)
// space: O(1)