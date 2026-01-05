func sumFourDivisors(nums []int) int {
    memo := map[int]int{}
    ans := 0
    for _, num := range nums {
        if val, exists := memo[num]; exists {
            ans += val
            continue
        }

        countDiv := 0
        divSum := 0
        four := true
        for i := 1; i <= num; i++ {
            if num % i == 0 {
                countDiv++
                divSum += i
            }

            if countDiv > 4 {
                four = false
                break
            }
        }

        if four && countDiv == 4 {
            ans += divSum
            memo[num] = divSum
        }
    }

    return ans
}

// array, math
// time: O(nm)
// space: O(n)