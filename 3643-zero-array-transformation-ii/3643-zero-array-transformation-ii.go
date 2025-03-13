func minZeroArray(nums []int, queries [][]int) int {
    diff := make([]int, len(nums)+1)
    k := 0
    for i := 0; i < len(nums); i++ {
        for ; diff[i] < nums[i] && k < len(queries); k++ {
            start, end, val := queries[k][0], queries[k][1], queries[k][2]
            if end < i {
                continue
            }
            diff[max(start, i)] += val
            diff[end+1] -= val
        }
        if diff[i] < nums[i] {
            return -1
        }
        diff[i+1] += diff[i]
    }

    return k
}

// notice me senpai

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}