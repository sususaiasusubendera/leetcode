func missingInteger(nums []int) int {
    set := map[int]bool{}
    for i := 0; i < len(nums); i++ {
       set[nums[i]] = true 
    }

    sum := nums[0]
    for i := 1; i < len(nums); i++ {
        if nums[i] != nums[i-1] + 1 {
            // break because a "prefix" is nums[0..i]
            // nums[n..i] with n > 0 isn't a valid "prefix"
            break 
        }

        sum += nums[i]
    }

    for set[sum] {
        sum++
    }

    return sum
}

// array, hash map
// time: O(n)
// space: O(n)

// n.b.
// the problem desc's suck
// a prefix is literally nums[0..i], FROM INDEX 0!!! FFFFFFF