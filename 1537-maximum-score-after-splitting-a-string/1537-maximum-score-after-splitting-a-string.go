func maxScore(s string) int {
    left, right, result := 0, 0, 0

    // precompute prefix sum of ones ('1')
    for i := 1; i < len(s); i++ {
        if s[i] == '1' {
            right++
        }
    }

    for i := 0; i < len(s)-1; i++ {
        if s[i] == '0' {
            left++
        }
        result = max(result, left + right)

        if s[i+1] == '1' {
            right--
        }
    }

    return result
}

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}