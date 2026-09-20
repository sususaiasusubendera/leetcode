func totalNumbers(digits []int) int {
    // 3-digit even numbers
    seen := make([]bool, 1000)
    ans := 0
    for i := 0; i < len(digits); i++ {
        if digits[i] == 0 {
            continue
        }
        for j := 0; j < len(digits); j++ {
            if j == i {
                continue
            }
            for k := 0; k < len(digits); k++ {
                if k == i || k == j || digits[k]%2 != 0 {
                    continue
                }
                num := digits[i]*100 + digits[j]*10 + digits[k]
                if !seen[num] {
                    seen[num] = true
                    ans++
                }
            }
        }
    }

    return ans
}

// array, brute force
// time: O(n^3)
// space: O(1)