func maxArea(height []int) int {
    left, right := 0, len(height)-1
    max := 0
    for left < right {
        if height[left] <= height[right] {
            area := height[left] * (right-left)
            if area > max { max = area }
            left++
        } else { // height[left] > height[right]
            area := height[right] * (right-left)
            if area > max { max = area }
            right--
        }
    }

    return max
}

// array, greedy, two pointers
// time: O(n)
// space: O(1)