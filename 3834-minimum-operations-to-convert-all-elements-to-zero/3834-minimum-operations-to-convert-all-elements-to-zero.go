func minOperations(nums []int) int {
    ops := 0
    stack := []int{}
    for _, num := range nums {
        for len(stack) > 0 && num < stack[len(stack)-1] {
            stack = stack[:len(stack)-1]
        }

        if num > 0 && (len(stack) == 0 || num > stack[len(stack)-1]) {
            ops++
            stack = append(stack, num)
        }
    }

    return ops
}

// array, monotonic stack
// time: O(n)
// space: O(n)