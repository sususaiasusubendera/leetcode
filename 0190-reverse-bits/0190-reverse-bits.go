func reverseBits(n int) int {
    bits := make([]int, 32)

    i := 0
    for n > 0 {
        bits[i] = n % 2
        n /= 2
        i++
    }

    ans := 0
    for i := 0; i < 32; i++ {
        ans += bits[i] * pow2(31-i)
    }

    return ans
}

func pow2(x int) int {
    res := 1
    for i := 0; i < x; i++ {
        res *= 2
    }
    return res
}

// notice me senpai