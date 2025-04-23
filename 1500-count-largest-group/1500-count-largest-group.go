func countLargestGroup(n int) int {
    m := map[int]int{}
    for i := 1; i <= n; i++ {
        m[digitSum(i)]++
    }

    count, maxFreq := 0, 0
    for _, freq := range m {
        if freq > maxFreq {
            maxFreq = freq
        }
    }

    for _, freq := range m {
        if freq == maxFreq {
            count++
        }
    }

    return count
}

// time: O(nlog(n))
// space: log(n)

func digitSum(n int) int {
    sum := 0
    for n != 0 {
        sum += n % 10
        n /= 10
    }

    return sum
}