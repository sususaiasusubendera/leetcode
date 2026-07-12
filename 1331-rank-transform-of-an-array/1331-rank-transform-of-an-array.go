func arrayRankTransform(arr []int) []int {
    sorted := make([]int, len(arr))
    copy(sorted, arr)
    sort.Ints(sorted)

    m := map[int]int{}
    visited := map[int]bool{}
    rank := 1
    for i := 0; i < len(sorted); i++ {
        if !visited[sorted[i]] {
            m[sorted[i]] = rank
            visited[sorted[i]] = true
            rank++
        }
    }

    ans := make([]int, len(arr))
    for i := 0; i < len(arr); i++ {
        ans[i] = m[arr[i]]
    }

    return ans
}

// array, hash map, sorting
// time: O(nlog(n))
// space: O(n)