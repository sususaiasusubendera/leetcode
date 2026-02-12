func longestBalanced(s string) int {
	ans := 0
	for i := 0; i < len(s); i++ {
		set := map[byte]int{}
		for j := i; j < len(s); j++ {
			set[s[j]]++

			balance := true
			currFreq := set[s[j]]
			for _, v := range set {
				if currFreq != v {
					balance = false
					break
				}
			}

			if balance {
				ans = max(ans, j-i+1)
			}
		}
	}

	return ans
}

// hash map, string
// time: O(n^2)
// space: O(1)