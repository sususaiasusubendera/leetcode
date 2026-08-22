func checkDivisibility(n int) bool {
    sum, prod := findSumProduct(n)

    return n % (sum + prod) == 0
}

func findSumProduct(n int) (int, int) {
    sum, prod := 0, 1
    for n > 0 {
        d := n % 10
        sum += d
        prod *= d
        n /= 10
    }

    return sum, prod
}

// math
// time: O(k)
// space: O(1)