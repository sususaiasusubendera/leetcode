func myPow(x float64, n int) float64 {
    if n == 0 {
        return 1.0
    }

    if n < 0 {
        x = 1 / x
        n = -n
    }

    result := 1.0
    currProd := x
    for n > 0 {
        if n % 2 == 1 {
            result *= currProd
        }
        currProd *= currProd
        n /= 2
    }
    return result
}