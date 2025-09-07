func sumZero(n int) []int {
    ans := []int{}
    if n % 2 == 1 {
        ans = append(ans, 0)
    }

    halfN := n/2
    counter := 1
    for halfN > 0 {
        ans = append(ans, counter)
        ans = append(ans, -counter)
        halfN--
        counter++
    }

    return ans
}

// time: O(n)
// space: O(n)