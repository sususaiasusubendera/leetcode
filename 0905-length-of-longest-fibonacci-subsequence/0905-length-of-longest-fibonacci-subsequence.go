func lenLongestFibSubseq(arr []int) int {
    ans := 0
    key, dp := map[int]int{}, map[[2]int]int{}
    
    for i, v := range arr {
        key[v] = i
        for j := i-1; j > 0 ; j-- {
            if arr[j] > arr[i]/2 {
                if v, exists := key[arr[i]-arr[j]]; exists {
                    dp[[2]int{i, j}] = dp[[2]int{j, v}] + 1
                    ans = max(ans, dp[[2]int{i, j}])
                }
            } else { break }
        }
    }
    if ans == 0 {
        return 0
    }
    return ans + 2
}

// NOTICE ME SENPAI!!!