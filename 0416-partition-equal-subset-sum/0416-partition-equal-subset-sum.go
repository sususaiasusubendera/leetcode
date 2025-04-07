func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	if sum%2 == 1 {
		return false
	}

	sum = sum / 2
    dp := make([]bool, sum)
    dp[0] = true 
    for _, n := range nums {
        if n <= sum {
            if dp[sum-n] == true {
                return true
            }

            for j := sum-n-1; j >= 0; j-- {
                if dp[j] == true {
                    dp[j+n] = true
                }
            }
        }
    }

    return false
}

// NOTICE ME SENPAI!!!