func minSubarray(nums []int, p int) int {
    sumMod := 0 // stores sum of nums % p, so that it doesn't overflow
    for _, num := range nums {
        sumMod = (sumMod + num) % p
    }

    need := sumMod
    if need == 0 { return 0 }

    m := map[int]int{} // map prefix[i] (in mod p) to i
    m[0] = 0 // prefix[0] = 0 at index 0
    res := len(nums) // length of the smallest subarray that need to remove
    prefix := 0
    for i := 1; i <= len(nums); i++ {
        prefix = (prefix + nums[i-1]) % p
        curr := prefix

        // let sum of subarray[l..r+1] = sub, and prefix sum = P
        // sub = P[r+1] - P[l]; sub % p == need
        // (P[r+1] - P[l]) % p == need) -> ((curr - P[l]) % p == need)
        // P[l] = (curr - need + p) % p; + p so that P[l] always positive [0..p)
        target := (curr - need + p) % p
        if idx, ok := m[target]; ok {
            // i - idx = length of subarray[idx..i]
            if i - idx < res { res = i - idx }
        }

        m[curr] = i
    }

    if res == len(nums) { return -1 }
    return res
}

// array, hash map, prefix sum 
// time: O(n)
// space: O(n)