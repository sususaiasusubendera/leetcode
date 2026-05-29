func minElement(nums []int) int {
    ans := make([]int, len(nums))
    for i := 0; i < len(nums); i++ {
        ans[i] = countSumDigit(nums[i])
    }

    minNum := 50
    for _, num := range ans {
        minNum = min(minNum, num)
    }

    return minNum
}

func countSumDigit(n int) int {
    sum := 0
    for n > 0 {
        sum += n % 10
        n /= 10
    }
    return sum
}

// array
// time: O(nd)
// space: O(1)