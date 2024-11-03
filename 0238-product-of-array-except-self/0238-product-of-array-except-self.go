func productExceptSelf(nums []int) []int {
    r := make([]int, len(nums))

    r[0] = 1
    for i := 1; i < len(nums); i++ {
        r[i] = r[i-1] * nums[i-1]
    }

    suffix := 1
    for i := len(nums)-1; i >= 0; i-- {
        r[i] *= suffix
        suffix *= nums[i]
    }

    return r

    // Time Limit Exceeded
    // var new []int
    // for idx, _ := range nums {
    //     v := 1
    //     for i := 0; i < len(nums); i++ {
    //         if i == idx {
    //             continue
    //         } else {
    //             v *= nums[i]
    //         }
    //     }
    //     new = append(new, v)
    // }
    // return new
}