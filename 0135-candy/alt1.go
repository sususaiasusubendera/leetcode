// greedy single-pass
func candy(ratings []int) int {
    candies := 1
    up, down, peak := 0, 0, 0
    for i := 1; i < len(ratings); i++ {
        if ratings[i] > ratings[i-1] {
            up++
            down = 0
            candies += up + 1
            peak = up
        } else if ratings[i] < ratings[i-1] {
            down++
            up = 0
            candies += down
            if peak < down {
                candies++
            }
        } else {
            up, down, peak = 0, 0, 0
            candies++
        }
    }
    return candies
}
