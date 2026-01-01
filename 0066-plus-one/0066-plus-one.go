func plusOne(digits []int) []int {
    idx := len(digits)-1
    for idx >= 0 && digits[idx] + 1 == 10 {
        digits[idx] = 0
        idx--
    }

    if idx == -1 {
        digits = append([]int{1}, digits...)
    } else {
        digits[idx]++
    }

    return digits
}

// array, math
// time: O(n)
// space: O(n)