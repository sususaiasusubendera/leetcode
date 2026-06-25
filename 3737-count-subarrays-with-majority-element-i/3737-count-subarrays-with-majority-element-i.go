func countMajoritySubarrays(nums []int, target int) int {
    ans := 0
    for i := 0; i < len(nums); i++ {
        count := 0
        for j := i; j < len(nums); j++ {
            if nums[j] == target {
                count++
            }

            if count > (j - i + 1) / 2 {
                ans++
            }
        }
    }

    return ans
}

// array, brute force
// time: O(n^2)
// space: O(1)

// try to get time O(n) (using prefix sum cmiiw)