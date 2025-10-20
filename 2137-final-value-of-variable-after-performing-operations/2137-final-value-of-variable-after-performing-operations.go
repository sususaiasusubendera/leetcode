func finalValueAfterOperations(operations []string) int {
    ans := 0
    for _, op := range operations {
        if op == "++X" || op == "X++" {
            ans++
            continue
        }

        ans--
    }

    return ans
}

// array, string
// time: O(n)
// space: O(1)