func punishmentNumber(n int) int {
    total := 0
    for i := 1; i <= n; i++ {
        square := i * i
        strSquare := strconv.Itoa(square)
        if canPartition(strSquare, i, 0, 0) {
            total += square
        }
    }
    return total
}

func canPartition(s string, target, sum, start int) bool {
    if start == len(s) {
        return sum == target
    }

    for end := start; end < len(s); end++ {
        num, _ := strconv.Atoi(s[start:end+1])
        if sum+num > target {
            break
        }
        if canPartition(s, target, sum+num, end+1) {
            return true
        }
    }

    return false
}