func maxDistance(nums1 []int, nums2 []int) int {
    ans := 0
    for i := 0; i < len(nums1); i++ {
        for j := i + ans; j < len(nums2); j++ {
            if nums1[i] > nums2[j] {
                break
            }

            ans = max(ans, j - i)
        }
    }

    return ans
}

// array
// time: O(nm)
// space: O(1)