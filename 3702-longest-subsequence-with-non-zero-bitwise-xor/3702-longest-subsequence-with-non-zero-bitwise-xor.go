func longestSubsequence(nums []int) int {
    totalXor := 0
    allZero := true
    for _, num := range nums {
        totalXor ^= num
        if num > 0 {
            allZero = false
        }
    }

    if totalXor > 0 {
        return len(nums)
    }

    if allZero {
        return 0
    }

    return len(nums) - 1
}

// array, bit manipulation
// time: O(n)
// space: O(1)

// n.b
// so, here the pattern
// first, know some xor properties: a ^ 0 = a, a ^ a = 0, a ^ b = b ^ a, a ^ (b ^ c) = (a ^ b) ^ c
// if the total xor is not 0, basically the longest subseq is the nums's length
// here's the magic
// if the total xor is 0, if there exists a minimum ONE num, with num > 0, the max length is the num's length - 1
// else, the max length is 0
// it's hard to believe at first
// try to simulate it on paper, and then the pattern makes sense