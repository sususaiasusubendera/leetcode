func totalWaviness(num1 int, num2 int) int {
    ans := 0
    for i := num1; i <= num2; i++ {
        ans += solve(i)
    }

    return ans
}

func solve(n int) int {
    s := strconv.Itoa(n)

    if len(s) < 3 {
        return 0
    }
    
    waviness := 0
    for i := 1; i < len(s)-1; i++ {
        if (s[i-1] < s[i] && s[i] > s[i+1]) || (s[i-1] > s[i] && s[i] < s[i+1]) {
            waviness++
        }
    }

    return waviness
}

// brute force
// time: O((num2 - num1) * log(num2))
// space: O(log(num2))