func isPalindrome(x int) bool {
    if x < 0 || (x != 0 && x % 10 == 0) {
        return false
    }

    reversedHalf := 0
    for x > reversedHalf {
        temp := x % 10
        reversedHalf = reversedHalf * 10 + temp
        x /= 10
    }

    return x == reversedHalf || x == reversedHalf / 10
}