func minSum(nums1 []int, nums2 []int) int64 {
	sum1, sum2 := 0, 0
    countZero1, countZero2 := 0, 0
    for _, num := range nums1 {
        sum1 += num
        if num == 0 {
            sum1++
            countZero1++
        }
    }

    for _, num := range nums2 {
        sum2 += num
        if num == 0 {
            sum2++
            countZero2++
        }
    }

    if ((sum1 < sum2) && countZero1 == 0) || ((sum2 < sum1) && countZero2 == 0) {
        return -1
    }

    if sum1 > sum2 {
        return int64(sum1)
    }

    return int64(sum2)
}

// greedy
// time: O(n + m)
// space: O(1)