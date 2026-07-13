func sequentialDigits(low int, high int) []int {
    secDigits := []int{}
    base := 1
    one := 1 // digits contain only num one
    for i := 2; i <= 9; i++ {
        one = (one * 10) + 1
        base += one
        n := base
        for j := 0; j < 10-i; j++ {
            secDigits = append(secDigits, n)
            n += one
        }
    }

    ans := []int{}
    for _, sd := range secDigits {
        if low <= sd && sd <= high {
            ans = append(ans, sd)
        }
    }

    return ans
}

// enumeration
// time: O(k)
// space: O(k)

// this solution and approach emerged from scribbles in notepad
// 2: 12 23 34 45 56 67 78 89 (8)
// 3: 123 234 345 456 567 678 789 (7)
// 4: 1234 2345 3456 4567 5678 6789 (6)
// 5: 12345 23456 34567 45678 56789 (5)
// 6: 123456 234567 345678 456789 (4)
// 7: 1234567 2345678 3456789 (3)
// 8: 12345678 23456789 (2)
// 9: 123456789 (1)