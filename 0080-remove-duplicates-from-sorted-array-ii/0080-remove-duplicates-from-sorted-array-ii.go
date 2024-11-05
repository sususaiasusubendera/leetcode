func removeDuplicates(nums []int) int {
    count := 1
    p := 1

    for i := 1; i < len(nums); i++ {
        if nums[i] == nums[i-1] {
            count++
            if count <= 2 {
                nums[p] = nums[i]
                p++
            }
        } else {
            count = 1
            nums[p] = nums[i]
            p++
        }
    }

    return p
}