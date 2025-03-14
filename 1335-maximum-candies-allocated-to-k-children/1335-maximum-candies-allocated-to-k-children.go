func maximumCandies(candies []int, k int64) int {
    left, right := 1, 0
    for _, pile := range candies {
        if pile > right {
            right = pile
        }
    }

    result := 0
    for left <= right {
        mid := left + (right-left)/2
        if canAllocate(candies, int(k), mid) {
            result = mid
            left = mid + 1
        } else {
            right = mid - 1
        }
    }

    return result
}

// NOTICE ME SENPAI!!!

func canAllocate(candies []int, k int, mid int) bool {
    count := 0
    for _, pile := range candies {
        count += pile / mid
        if count >= k {
            return true
        }
    }

    return false
}