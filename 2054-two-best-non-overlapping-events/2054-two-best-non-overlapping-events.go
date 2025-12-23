func maxTwoEvents(events [][]int) int {
    sort.Slice(events, func(i, j int) bool {
        return events[i][0] < events[j][0]
    })

    starts := [][]int{}
    for _, e := range events {
        if len(starts) == 0 || starts[len(starts)-1][0] < e[0] {
            starts = append(starts, []int{e[0], e[2]})
        } else {
            starts[len(starts)-1][1] = max(starts[len(starts)-1][1], e[2])
        }
    }

    for i := len(starts)-2; i >= 0; i-- {
        starts[i][1] = max(starts[i][1], starts[i+1][1])
    }

    out := 0
    for _, e := range events {
        end, val := e[1], e[2]
        index, _ := slices.BinarySearchFunc(starts, []int{end+1, 69}, func(a, b []int) int {
            return a[0] - b[0]
        })
        if index == len(starts) || starts[index][0] <= end {
            out = max(out, val)
        } else {
            out = max(out, val + starts[index][1])
        }
    }

    return out
}

// bin search, prefix sum, sorting
// solution from solutions @pattobears
// NOTICE ME SENPAI