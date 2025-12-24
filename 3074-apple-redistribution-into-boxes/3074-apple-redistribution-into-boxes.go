func minimumBoxes(apple []int, capacity []int) int {
    sort.Slice(capacity, func(i, j int) bool {
        return capacity[i] > capacity[j]
    })

    totalApple := 0
    for _, a := range apple {
        totalApple += a
    }

    ans := 0
    idxc := 0
    for totalApple > 0 && idxc < len(capacity) {
        totalApple -= capacity[idxc]
        ans++
        idxc++
    }

    return ans
}

// array, greedy, sorting
// time: O(nlog(n) + m)
// space: O(1)