func queryResults(limit int, queries [][]int) []int {
    result := make([]int, len(queries))
    mapBallColor := map[int]int{} // a ball -> color map
    mapColorCountBalls := map[int]int{} // a color -> count balls map
    countDistColors := 0 // count distinct colors each query
    for i, query := range queries {
        ball, newColor := query[0], query[1]

        if prevColor, exists := mapBallColor[ball]; exists {
            mapColorCountBalls[prevColor]--
            // -1 countDistColors if prevColor no longer exists in this query
            if mapColorCountBalls[prevColor] == 0 {
                countDistColors--
            }
        }

        mapBallColor[ball] = newColor
        // +1 countDistColors if there is a new distinct color in this query
        if mapColorCountBalls[newColor] == 0 {
            countDistColors++
        }
        mapColorCountBalls[newColor]++

        result[i] = countDistColors
    }

    return result
}

// two hashmaps
// time: O(n)
// space: O(m + n)