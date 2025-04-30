func findNumbers(nums []int) int {
    res := 0

    for _, n := range nums {
        if (n > 9 && n < 100) || (n > 999 && n < 10000) || n == 100000 {
            res++
        }
    }

    return res    
}

// notice me senpai