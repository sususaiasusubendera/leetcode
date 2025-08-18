func judgePoint24(cards []int) bool {
    floatCards := make([]float64, 4)
    for i, card := range cards {
        floatCards[i] = float64(card)
    }

    var solve func(nums []float64) bool
    solve = func(nums []float64) bool {
        if len(nums) == 1 {
            return isEqual(nums[0], 24.0)
        }

        for i := 0; i < len(nums); i++ {
            for j := 0; j < len(nums); j++ {
                if i == j {
                    continue
                }

                a, b := nums[i], nums[j]
                res := make([]float64, 6)
                res[0] = a + b
                res[1] = a - b
                res[2] = b - a
                res[3] = a * b
                res[4] = a / b // TODO: check if b != 0
                res[5] = b / a // TODO: check if a != 0
                for k := 0; k < len(res); k++ {
                    next := []float64{}
                    next = append(next, res[k])
                    for l := 0; l < len(nums); l++ {
                        if l != i && l != j {
                            next = append(next, nums[l])
                        }
                    }
                    
                    if solve(next) {
                        return true
                    }
                }
            }
        }

        return false
    }

    return solve(floatCards)
}

// backtracking
// time: O((n!)^2); O(1) (leetcode)
// space: O(n^2); O(1) (leetcode)

func isEqual(a, b float64) bool {
    const eps = 1e-6
    return abs(a-b) < eps
}

func abs(n float64) float64 {
    if n < 0 {
        return -n
    }

    return n
}