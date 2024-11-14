func isHappy(n int) bool {
    count := 0

    happy := false
    for !happy {
        temp := 0
        for n > 0 {
            temp += pow(n % 10, 2)
            n /= 10
        }
        count++
        if temp == 1 {
            return !happy
        } else {
            n = temp
        }
        if count > 100 {
            return happy
        }
    }

    return happy
}

func pow(n, p int) int {
    a := 1
    for i := 0; i < p; i++ {
        a *= n
    }
    return a
}
