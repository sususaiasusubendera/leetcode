func candy(ratings []int) int {
    candies := make([]int, len(ratings))
    for i := 0; i < len(candies); i++ {
        candies[i] = 1
    }

    // check from left to right
    for i := 1; i < len(ratings); i++ {
        if ratings[i] > ratings[i-1] {
            candies[i] = candies[i-1] + 1
        }
    }

    // check from right to left
    for i := len(ratings)-2; i >= 0; i-- {
        if (ratings[i] > ratings[i+1]) && (candies[i] <= candies[i+1]) {
            // add a candy if candies[i] is not more than candies[i+1]
            candies[i] = candies[i+1] + 1
        }
    }

    result := 0
    for _, candy := range candies {
        result += candy
    }

    return result
}

// greedy two-pass
// time: O(n)
// space: O(n)

// DCC june 2025 day 2