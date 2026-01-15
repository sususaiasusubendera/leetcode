func maximizeSquareHoleArea(n int, m int, hBars []int, vBars []int) int {
	sort.Slice(hBars, func(i, j int) bool { return hBars[i] < hBars[j] })
	sort.Slice(vBars, func(i, j int) bool { return vBars[i] < vBars[j] })

    hMax, vMax := 1, 1 // longest h and v
	h, v := 1, 1
    if len(hBars) > 1 {
        for i := 1; i < len(hBars); i++ {
            if hBars[i] == hBars[i-1] + 1 {
                h++
                hMax = max(hMax, h)
            } else {
                h = 1
            }
        }
    }
    if len(vBars) > 1 {
        for i := 1; i < len(vBars); i++ {
            if vBars[i] == vBars[i-1] + 1 {
                v++
                vMax = max(vMax, v)
            } else {
                v = 1
            }
        }
    }

    s := min(hMax, vMax) + 1 // side
    return s * s
}

// array, sorting
// time:(O(nlogn + mlogm))
// space: O(1)