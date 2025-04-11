func countSymmetricIntegers(low int, high int) int {
    count := 0
    for i := low; i <= high; i++ {
        if i < 100 && i % 11 == 0 {
            count++
        }

        if 1000 <= i && i < 10000 {
            left := (i/1000) + ((i%1000)/100)
            right := ((i%100)/10) + (i%10)
            if left == right {
                count++
            }
        }
    }

    return count
}

// time: O(n)
// space: O(1)