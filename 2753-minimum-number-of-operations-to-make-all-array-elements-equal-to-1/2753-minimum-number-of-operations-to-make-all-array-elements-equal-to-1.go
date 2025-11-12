func minOperations(nums []int) int {
	countOnes := 0
	for _, num := range nums {
		if num == 1 { countOnes++ }
	}
	if countOnes > 0 { return len(nums) - countOnes }

	minLen := len(nums) + 1
	for i := 0; i < len(nums); i++ {
		g := nums[i]
		for j := i + 1; j < len(nums); j++ {
			g = gcd(g, nums[j])
			if g == 1 {
				minLen = min(minLen, j-i+1)
				break // because any longer subarray will also give gcd = 1
			}
		}
	}

	// no subarray can produce gcd = 1
	if minLen == len(nums)+1 { return -1 }

	// minLen - 1: ops to create one a single 1 in the shortest subarray
	// len(nums) - 1: ops to spread that 1 to all elements
	return (minLen - 1) + (len(nums) - 1)
}

// array, math
// time: O(n^2)
// space: O(1)

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	if a < 0 { return -a }
	return a
}

func min(a, b int) int {
	if a < b { return a }
	return b
}