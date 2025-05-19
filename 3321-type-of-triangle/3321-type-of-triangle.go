func triangleType(nums []int) string {
    if nums[0] == nums[1] && nums[1] == nums[2] && nums[0] == nums[2] {
        return "equilateral"
    }

    if (nums[0] + nums[1] > nums[2]) && (nums[0] + nums[2] > nums[1]) && (nums[1] + nums[2] > nums[0]) {
        if nums[0] == nums[1] || nums[1] == nums[2] || nums[0] == nums[2] {
            return "isosceles"
        }

        return "scalene"
    }

    return "none"
}

// time: O(1)
// space: O(1)