func totalFruit(fruits []int) int {
    basket := map[int]int{}
    left := 0
    maxFruits := 0
    for right := 0; right < len(fruits); right++ {
        basket[fruits[right]]++

        for len(basket) > 2 {
            basket[fruits[left]]--
            if basket[fruits[left]] == 0 {
                delete(basket, fruits[left])
            }
            left++
        }

        countFruits := right - left + 1
        if countFruits > maxFruits {
            maxFruits = countFruits
        }
    }

    return maxFruits
}

// hash map, sliding window
// time: O(n)
// space: O(1)