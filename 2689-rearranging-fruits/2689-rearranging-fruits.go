func minCost(basket1 []int, basket2 []int) int64 {
	freq := map[int]int{}
    minCost := int(1e9) // from constraints 1 <= basket1[i], basket2[i] <= 10^9
    for _, cost := range basket1 {
        freq[cost]++
        if cost < minCost {
            minCost = cost
        }
    }
    for _, cost := range basket2 {
        freq[cost]--
        if cost < minCost {
            minCost = cost
        }
    }

    swap := []int{}
    for k, v := range freq {
        if v % 2 != 0 {
            return -1
        }

        for i := 0; i < abs(v)/2; i++ {
            swap = append(swap, k)
        }
    }

    sort.Ints(swap)
    result := 0
    for i := 0; i < len(swap)/2; i++ {
        if 2*minCost < swap[i] {
            result += 2*minCost
            continue
        }

        result += swap[i]
    }

    return int64(result)
}

// hash map, greedy, sort
// time: O(Nlog(n))
// space: O(N)

func abs(n int) int {
    if n < 0 {
        return -n
    }

    return n
}