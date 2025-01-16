func xorAllNums(nums1 []int, nums2 []int) int {
    x1, x2 := 0, 0

    // if len of nums1 is odd
    if len(nums1) % 2 != 0 {
        for _, num := range nums2 {
            x2 ^= num
        }
    }

    // if len of nums2 is odd
    if len(nums2) % 2 != 0 {
        for _, num := range nums1 {
            x1 ^= num
        }
    }

    return x1 ^ x2
}

// xor property approach
// time: O(m + n)
// space: O(1)

// p.s. xor properties
// a xor b = b xor a
// (a xor b) xor c = a xor (b xor c)
// a xor a = 0
// a xor 0 = a