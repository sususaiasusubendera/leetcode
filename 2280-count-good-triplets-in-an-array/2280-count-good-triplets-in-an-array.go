func goodTriplets(nums1 []int, nums2 []int) int64 {
    n := len(nums1)
    pos2, reversedIndexMapping := make([]int, n), make([]int, n)
    for i, num := range nums2 {
        pos2[num] = i
    }

    for i, num := range nums1 {
        reversedIndexMapping[pos2[num]] = i
    }

    tree := newFenwickTree(n)
    var res int64
    for v := 0; v < n; v++ {
        pos := reversedIndexMapping[v]
        left := tree.query(pos)
        tree.update(pos, 1)
        right := (n - 1 - pos) - (v - left)
        res += int64(left * right)
    }

    return res
}

// NOTICE ME SENPAIII!!!!!!!

type FenwickTree struct {
    tree []int
}

func newFenwickTree(n int) *FenwickTree {
    return &FenwickTree{
        tree: make([]int, n+1),
    }
}

func (ft *FenwickTree) update(i, delta int) {
    i++ // convert to 1-based index
    for i < len(ft.tree) {
        ft.tree[i] += delta
        i += i & -i
    }
}

func (ft *FenwickTree) query(i int) int {
    i++
    res := 0
    for i > 0 {
        res += ft.tree[i]
        i -= i & -i
    }

    return res
}