func countValidSelections(nums []int) int {
    countZero := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] == 0 { countZero++ }
    }

    countValid := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] == 0 {
            // direction right (1)
            if isValid(countZero, i, 1, nums) { countValid++ }

            // direction left (-1)
            if isValid(countZero, i, -1, nums) { countValid++ }
        }
    }

    return countValid
}

func isValid(countZero, idx, dir int, nums []int) bool {
    temp := make([]int, len(nums))
    copy(temp, nums)
    curr := idx
    for 0 <= curr && curr <= len(temp)-1 {
        if temp[curr] > 0 {
            temp[curr]--
            if temp[curr] == 0 { countZero++ }
            dir *= -1 // reverse
        }
        curr += dir
    }

    if countZero == len(temp) { return true }
    return false
}

// array, simulation
// time: O(n^3)
// space: O(n)