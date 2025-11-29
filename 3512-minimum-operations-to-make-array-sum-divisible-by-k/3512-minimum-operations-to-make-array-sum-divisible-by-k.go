func minOperations(nums []int, k int) int {
    sum := 0
    for _, num := range nums {
        sum += num
    }

    return sum % k
}

// array, math
// time: O(n)
// space: O(1)