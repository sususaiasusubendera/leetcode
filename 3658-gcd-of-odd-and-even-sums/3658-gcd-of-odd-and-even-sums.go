func gcdOfOddEvenSums(n int) int {
    sumOdd, sumEven := 0, 0
    for i := 1; i <= 2*n; i++ {
        if i % 2 == 0 {
            sumEven += i
            continue
        }
        sumOdd += i
    }

    return gcd(sumOdd, sumEven)
}

func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a % b
    }
    return a
}

// math
// time: O(n)
// space: O(1)