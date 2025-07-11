func differenceOfSums(n int, m int) int {
    num1, num2 := 0, 0
    for i := 1; i <= n; i++ {
        if i % m != 0 {
            num1 += i
            continue
        }

        num2 += i
    }

    return num1 - num2
}

// time: O(n)
// space: O(1)