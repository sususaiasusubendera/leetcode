func longestNiceSubarray(nums []int) int {
  usedBits, windowStart, maxLength := 0, 0, 0
  for windowEnd := 0; windowEnd < len(nums); windowEnd++ {
    for (usedBits & nums[windowEnd]) != 0 {
        usedBits ^= nums[windowStart]
        windowStart++
    }

    usedBits |= nums[windowEnd]

    maxLength = max(maxLength, windowEnd - windowStart + 1)
  }

  return maxLength
}

// NOTICE ME SENPAI!!!

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}