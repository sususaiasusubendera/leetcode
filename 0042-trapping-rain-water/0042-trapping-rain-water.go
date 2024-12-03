func trap(height []int) int {
    if len(height) <= 2 {
        return 0
    }

    left, right := 0, len(height)-1
    maxLeft, maxRight := height[left], height[right]
    water := 0
    for left < right {
        if height[left] < height[right] {
            if maxLeft > height[left] {
                water += maxLeft - height[left]
            }
            maxLeft = max(maxLeft, height[left])
            left++
        } else {
            if maxRight > height[right] {
                water += maxRight - height[right]
            }
            maxRight = max(maxRight, height[right])
            right--
        }
    }
    return water
}

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func min(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}