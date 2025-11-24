func prefixesDivBy5(nums []int) []bool {
    ans := make([]bool, len(nums))
    curr := 0
    for i, bit := range nums {
        curr = (curr*2 + bit) % 5
        ans[i] = curr == 0
    }

    return ans
}

// array, bit manipulation
// time: O(n)
// space: O(n)