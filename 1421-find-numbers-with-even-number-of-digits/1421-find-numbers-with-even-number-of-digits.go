func findNumbers(nums []int) int {
    count := 0
    for _, num := range nums {
        if (9 < num && num < 100) || (999 < num && num < 10000) || num == 100000 {
            count++
        }
    }

    return count
}

// time: O(n)
// space: O(1)