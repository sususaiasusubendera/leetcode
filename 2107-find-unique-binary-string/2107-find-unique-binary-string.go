func findDifferentBinaryString(nums []string) string {
    numSet := map[string]bool{}
    for _, num := range nums {
        numSet[num] = true
    }

    n := len(nums)

    var generate func(string) string
    generate = func(curr string) string {
        if len(curr) == n {
            if !numSet[curr] {
                return curr
            }
            return ""
        }

        addZero := generate(curr + "0")
        if addZero != "" {
            return addZero
        }

        return generate(curr + "1")
    }

    return generate("")
}

// notice me senpai