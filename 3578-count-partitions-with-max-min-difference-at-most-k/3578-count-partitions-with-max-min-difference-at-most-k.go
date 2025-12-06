func countPartitions(nums []int, k int) int {
    const MOD = 1_000_000_007

    n := len(nums)
    if n == 0 { return 0 }

    dp := make([]int, n+1) // dp[i]: number of valid partitions for nums[0..i-1]
    pre := make([]int, n+2) // prefix sum: pre[i] = sum(dp[0..i-1])
    
    // mono deque, storing indices of nums[]
    maxD := []int{} // indices with decreasing num (front: max)
    minD := []int{} // indices with increasing num (front: min)
    
    dp[0] = 1 
    pre[0] = 0
    pre[1] = dp[0]
    left := 0
    for i := 1; i <= n; i++ {
        idx := i - 1 // convert dp-index i -> nums-index i

        // update mono max deque
        for len(maxD) > 0 && nums[maxD[len(maxD)-1]] <= nums[idx] { maxD = maxD[:len(maxD)-1] } // pop back
        maxD = append(maxD, idx)

        // update mono min deque
        for len(minD) > 0 && nums[minD[len(minD)-1]] >= nums[idx] { minD = minD[:len(minD)-1]} // pop back
        minD = append(minD, idx)

        // shift left until window nums[left..idx] is valid
        for {
            max := nums[maxD[0]]
            min := nums[minD[0]]

            if max - min <= k { break } // valid
            
            // remove indices leaving the window
            if maxD[0] == left { maxD = maxD[1:] } // pop front
            if minD[0] == left { minD = minD[1:] } // pop front
            left++
        }

        // dp[i] = sum(dp[left..i-1]) = pre[i] - pre[left]
        dp[i] = (pre[i] - pre[left] + MOD) % MOD

        // update prefix sum: pre[i+1] = sum(dp[0..i])
        pre[i+1] = (pre[i] + dp[i]) % MOD
    }

    return dp[n]
}

// maybe senpai can notice this
// array, dp (bot-up), mono dequeue, prefix sum, sliding window
// time: O(n)
// space: O(n)