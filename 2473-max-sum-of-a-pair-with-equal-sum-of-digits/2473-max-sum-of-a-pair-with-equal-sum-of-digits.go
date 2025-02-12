func maximumSum(nums []int) int {
    maxSum := -1
    mapSumDigitsIndexes := map[int][]int{} // grouping indexes by sum of digits
    for i := 0; i < len(nums); i++ {
        if nums[i] > 9 { // nums[i] that has > 1 digit
            num := nums[i]
            sumDigits := 0
            for num != 0 {
                digit := num % 10
                num /= 10
                sumDigits += digit
            }
            mapSumDigitsIndexes[sumDigits] = append(mapSumDigitsIndexes[sumDigits], i)
        } else { // nums[i] that has 1 digit
            mapSumDigitsIndexes[nums[i]] = append(mapSumDigitsIndexes[nums[i]], i)
        }
    }

    for _, v := range mapSumDigitsIndexes {
        // find the largest sum of 2 elements of each group (mapSumDigitsIndexes)
        if len(v) > 1 {
            for i := 0; i < len(v)-1; i++ {
                for j := i+1; j < len(v); j++ {
                    sum := nums[v[i]] + nums[v[j]]
                    if sum > maxSum {
                        maxSum = sum
                    }
                }
            }
        }
    }

    return maxSum
}

// time: O(n.log(m) + n^2)
// space: O(n)