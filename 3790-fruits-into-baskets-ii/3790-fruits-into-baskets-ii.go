func numOfUnplacedFruits(fruits []int, baskets []int) int {
    unplaced := len(fruits)
    used := make([]bool, len(baskets))
    for i := 0; i < len(fruits); i++ {
        for j := 0; j < len(baskets); j++ {
            if !used[j] && fruits[i] <= baskets[j] {
                unplaced--
                used[j] = true
                break
            }
        }
    }

    return unplaced
}

// brute force
// time: O(nm)
// space: O(m)