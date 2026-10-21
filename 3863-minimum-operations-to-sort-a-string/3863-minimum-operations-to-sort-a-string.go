func minOperations(s string) int {
    minChar, maxChar := byte('z'), byte('a')
    sorted := true
    for i := 0; i < len(s); i++ {
        if s[i] < minChar {
            minChar = s[i]
        }

        if s[i] > maxChar {
            maxChar = s[i]
        }

        if i > 0 && s[i-1] > s[i] {
            sorted = false
        }
    }

    if sorted {
        return 0
    }

    if len(s) == 2 {
        return -1
    }

    if s[0] == minChar || s[len(s)-1] == maxChar {
        return 1
    }

    for i := 1; i < len(s)-1; i++ {
        if s[i] == minChar || s[i] == maxChar {
            return 2
        }
    }

    // s[0] == maxChar && s[len(s)-1] == minChar
    return 3
}

// string
// time: O(n)
// space: O(1)