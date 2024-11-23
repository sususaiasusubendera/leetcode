func hIndex(citations []int) int {
    citations = countingSort(citations)
    for i := 0; i < len(citations); i++ {
        if citations[i] < i + 1 {
            return i
        }
    }
    return len(citations)
}

func countingSort(a []int) []int {
    if len(a) == 0 {
        return a
    }

    min, max := a[0], a[0]
    for _, v := range a {
        if v > max {
            max = v
        }
        if v < min {
            min = v
        }
    }

    count := make([]int, max - min + 1)
    for _, v := range a {
        count[v-min]++
    }

    idx := len(a)-1
    for i, c := range count {
        for c > 0 {
            a[idx] = i + min
            idx--
            c--
        }
    }

    return a
}