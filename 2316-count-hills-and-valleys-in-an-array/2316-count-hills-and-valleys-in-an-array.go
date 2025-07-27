func countHillValley(nums []int) int {
    newNums := []int{}
    newNums = append(newNums, nums[0])
    p := 1
    for p < len(nums) {
        for p < len(nums) && nums[p-1] == nums[p] {
            p++
        }

        if p < len(nums) {
            newNums = append(newNums, nums[p])
            p++
        }
    }

    result := 0
    for i := 0; i < len(newNums); i++ {
        if i == 0 || i == len(newNums)-1 {
            continue
        }

        if newNums[i-1] > newNums[i] && newNums[i] < newNums[i+1] { // valley
            result++
        } else if newNums[i-1] < newNums[i] && newNums[i] > newNums[i+1] { // hill
            result++
        }
    }

    return result
}

// time: O(n)
// space: O(n)