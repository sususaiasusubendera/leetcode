func findEvenNumbers(digits []int) []int {
    freq := [10]int{}
    for _, digit := range digits {
        freq[digit]++
    }

    result := []int{}
    for num := 100; num < 999; num += 2 {
        digitFreq := [10]int{}
        digitFreq[num%10]++
        digitFreq[(num/10)%10]++
        digitFreq[(num/100)%10]++

        found := true
        for i := 0; i < 10; i++ {
            if freq[i] < digitFreq[i] {
                found = false
                break
            }
        }

        if found {
            result = append(result, num)
        }
    }

    return result
}

// time: O(n)
// space: O(1)