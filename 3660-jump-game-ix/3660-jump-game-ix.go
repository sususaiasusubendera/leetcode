func maxValue(nums []int) []int {
    stack := []Item{} // monotonic increasing stack
    for i := 0; i < len(nums); i++ {
        curr := Item{nums[i], i, i}

        for len(stack) > 0 && stack[len(stack)-1].val > nums[i] {
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1] // pop
            if top.val > curr.val {
                curr.val = top.val
            }
            curr.left = top.left
        }

        stack = append(stack, curr)
    }

    ans := make([]int, len(nums))
    for i := 0; i < len(stack); i++ {
        val, left, right := stack[i].val, stack[i].left, stack[i].right
        for j := left; j <= right; j++ {
            ans[j] = val
        }
    }

    return ans
}

type Item struct {
	val   int // nums[i]
	left  int // left idx
	right int // right idx
}

// monotonic stack
// time: O(n)
// space: O(n)

// note
// i < j and nums[i] > nums[j] (i can jump to j, j can also jump to i, an undirected graph)
// maintain connected components with a monotonic stack
// smaller value merges all previous greater components