func intersectionSizeTwo(intervals [][]int) int {
    // sort by end asc; start desc
    sort.Slice(intervals, func(i, j int) bool {
        if intervals[i][1] == intervals[j][1] {
            return intervals[i][0] > intervals[j][0]
        }
        return intervals[i][1] < intervals[j][1]
    })

    ans := 0
    a, b := -1, -1 // a < b; current 2 largest chosen numbers
    for _, interval := range intervals {
        start, end := interval[0], interval[1]
        if start > b { // interval has no intersection with the two chosen
            ans += 2
            a = end - 1
            b = end
        } else if start > a { // interval intersects with b but not with a
            ans += 1
            a = b
            b = end
        } else {
            // start <= a: both a and b are inside interval -> nothing to do
            continue
        }
    }

    return ans
}

// array, greedy, sorting
// time: O(nlog(n))
// space: O(1)
