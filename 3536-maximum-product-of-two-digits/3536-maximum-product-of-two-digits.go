func maxProduct(n int) int {
    temp := []int{}
    m := 1
    for n > 0 {
        d := n % (m * 10)
        temp = append(temp, d)
        n /= 10
    }

    sort.Ints(temp)
    
    return temp[len(temp)-1] * temp[len(temp)-2]
}

// math, sorting
// time: O(klog(k))
// space: O(k)