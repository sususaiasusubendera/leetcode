func kidsWithCandies(candies []int, extraCandies int) []bool {
    mc := candies[0] // mc: most candies
    for i := 1; i < len(candies); i++ {
        if candies[i] > mc {
            mc = candies[i]
        }
    }

    result := make([]bool, len(candies))
    for i := 0; i < len(candies); i++ {
        result[i] = candies[i] + extraCandies >= mc
    }

    return result
}