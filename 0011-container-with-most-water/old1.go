func maxArea(height []int) int {
    x1 := 0
    x2 := len(height)-1

    ma := 0
    for x1 <= x2 {
        a := 0
        if height[x1] >= height[x2] {
            a = height[x2] * (x2 - x1)
            if a > ma {
                ma = a
            }
            x2--
        } else {
            a = height[x1] * (x2 - x1)
            if a > ma {
                ma = a
            }
            x1++
        }
    }

    return ma
}
