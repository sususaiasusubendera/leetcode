func maxPalindromes(s string, k int) int {
    ans, start := 0, 0
    for right := k - 1; right < len(s); right++ {
        // palindrome with size k
        left := right - k + 1
        if left >= start && isPal(left, right, s) {
            ans++
            start = right + 1
            continue
        }

        // palindrome with size k + 1
        left = right - k
        if left >= start && isPal(left, right, s) {
            ans++
            start = right + 1
        }
    }

    return ans    
}

func isPal(l, r int, s string) bool {
    for l < r {
        if s[l] != s[r] {
            return false
        }
        l++
        r--
    }

    return true 
}

// greedy (the greedy idea is from editorial)
// time: O(nk)
// space: O(1)