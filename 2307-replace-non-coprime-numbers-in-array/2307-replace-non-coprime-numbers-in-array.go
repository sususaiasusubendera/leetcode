func replaceNonCoprimes(nums []int) []int {
    if len(nums) == 1 { return nums }

    stack := []int{}
    for i := 0; i < len(nums); i++ {
        stack = append(stack, nums[i])
        for len(stack) > 1 && GCD(stack[len(stack)-2], stack[len(stack)-1]) > 1 {
            a, b := stack[len(stack)-2], stack[len(stack)-1]
            stack = stack[:len(stack)-2] // pop 2x
            stack = append(stack, LCM(a, b)) // push LCM(top2, top1)
        }
    }

    return stack
}

// array, math, stack
// time: O(n log(M))
// space: O(n)

func GCD(a, b int) int {
    for b != 0 {
        a, b = b, a%b
    }
    return a
}

func LCM(a, b int) int {
    return (a / GCD(a, b)) * b
}