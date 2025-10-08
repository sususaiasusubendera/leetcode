func successfulPairs(spells []int, potions []int, success int64) []int {
    sort.Ints(potions)
    res := []int{}
    for _, spell := range spells {
        minPotion := (int(success) + spell - 1) / spell // ceil div

        left, right := 0, len(potions)-1
        for left < right {
            mid := left + (right-left)/2
            if potions[mid] >= minPotion {
                right = mid
            } else {
                left = mid + 1
            }
        }

        if potions[left] < minPotion {
            res = append(res, 0)
        } else {
            res = append(res, len(potions)-left)
        }
    }

    return res
}

// binary search
// time: O((n + m) log(n))
// space: O(m)