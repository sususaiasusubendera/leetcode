func findLHS(nums []int) int {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}

	result := 0
	for num, count := range freq {
		if count2, ok := freq[num+1]; ok {
			result = max(result, count+count2)
		}
	}
	return result
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}