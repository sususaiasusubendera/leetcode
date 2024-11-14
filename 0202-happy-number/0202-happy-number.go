func isHappy(n int) bool {
    slow, fast := n, next(n)
    for fast != 1 && slow != fast {
        slow = next(slow)
        fast = next(next(fast))
    }
    return fast == 1
}

func next(n int) int {
    sum := 0
    for n > 0 {
        digit := n % 10
        sum += digit * digit
        n /= 10
    }
    return sum
}