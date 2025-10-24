func nextBeautifulNumber(n int) int {
    for i := n + 1; i <= 1224444; i++ { // constraint: 0 <= n <= 10^6
        if isBalance(i) { return i }
    }

    return -1
}

func isBalance(x int) bool {
    digitFreq := map[int]int{}
    for x > 0 {
        digitFreq[x%10]++
        x /= 10
    }

    for digit, freq := range digitFreq {
        if digit != freq { return false }
    }

    return true
}

// hash map
// time: O(1224444)
// space: O(1)