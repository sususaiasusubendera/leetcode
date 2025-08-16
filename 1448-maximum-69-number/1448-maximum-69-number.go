func maximum69Number (num int) int {
    div := 1
    for div <= num/10 {
        div *= 10
    }

    res := 0
    change := false
    for div > 0 {
        d := num / div
        if d == 6 && !change {
            d = 9
            change = true
        }
        num %= div
        res += d * div
        div /= 10
    }

    return res
}

// greedy, integer
// time: O(log(n))
// space: O(1)