func tupleSameProduct(nums []int) int {
    // find all tupples' product
    mapProd := map[int]int{}
    for i := 0; i < len(nums)-1; i++ {
        for j := i+1; j < len(nums); j++ {
            prod := nums[i] * nums[j]
            mapProd[prod]++
        }
    }

    // calculate all C(n, 2) for products that appear ≥ 2 times
    // C(n, 2), 2 unique pairs that have the same products
    comb := 0
    for _, v := range mapProd {
        if v > 1 {
            comb += combTwo(v)
        }
    }

    // multiply the combination result by 8
    return comb * 8
    // multiply by 8 because each pair (a, b) and (c, d) can be arranged in:
    // - 2 ways by swapping a and b
    // - 2 ways by swapping c and d
    // - 2 ways by swapping (a, b) with (c, d)
    // total: 2 × 2 × 2 = 8 (permutation)
}

// time: O(n^2)
// time: O(n)

// UTILS
// combination formula C(n, 2)
func combTwo(n int) int {
    return (n * (n-1)) / 2
}