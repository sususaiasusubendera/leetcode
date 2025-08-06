func numOfUnplacedFruits(fruits []int, baskets []int) int {
    n := len(baskets)
    if n == 0 {
        return n
    }

    st := SegmentTree{
        tree: make([]int, 4*n),
        baskets: baskets,
    }
    st.build(1, 0, n-1)

    count := 0
    for i := 0; i < len(fruits); i++ {
        l, r, res := 0, n-1, -1
        for l <= r {
            mid := (l + r) >> 1
            if st.query(1, 0, n-1, 0, mid) >= fruits[i] {
                res = mid
                r = mid - 1
            } else {
                l = mid + 1
            }
        }

        if res != -1 && st.baskets[res] >= fruits[i] {
            st.update(1, 0, n-1, res, -1)
        } else {
            count++
        }
    }

    return count
}

// segment tree
// time: O(n + mlog(n))
// space: O(n)

type SegmentTree struct {
    tree []int
    baskets []int
}

func (st *SegmentTree) build(node, l, r int) {
    if l == r {
        st.tree[node] = st.baskets[l]
        return
    }

    mid := (l + r) >> 1 // (l + r) / 1
    st.build(node<<1, l, mid) // left child (node*2)
    st.build(node<<1|1, mid+1, r) // right child (node*2+1)
    st.tree[node] = max(st.tree[node<<1], st.tree[node<<1|1])
}

func (st *SegmentTree) query(node, l, r, ql, qr int) int {
    if ql > r || qr < l {
        return -1
    }

    if ql <= l && r <= qr {
        return st.tree[node]
    }

    mid := (l + r) >> 1
    return max(st.query(node<<1, l, mid, ql, qr), st.query(node<<1|1, mid+1, r, ql, qr))
}

func (st *SegmentTree) update(node, l, r, pos, val int) {
    if l == r {
        st.tree[node] = val
        return
    }

    mid := (l + r) >> 1
    if pos <= mid {
        st.update(node<<1, l, mid, pos, val)
    } else {
        st.update(node<<1|1, mid+1, r, pos, val)
    }

    st.tree[node] = max(st.tree[node<<1], st.tree[node<<1|1])
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}