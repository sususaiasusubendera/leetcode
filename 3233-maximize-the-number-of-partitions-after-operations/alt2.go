func maxPartitionsAfterOperations(s string, k int) int {
    best := countPartitions(s, k)

    // change 1 char and try all the possibilities
    for i := 0; i < len(s); i++ {
        for c := 'a'; c <= 'z'; c++ {
            if byte(c) == s[i] { continue }

            temp := []byte(s)
            temp[i] = byte(c)
            count := countPartitions(string(temp), k)
            if count > best { best = count }
        }
    }

    return best
}

func countPartitions(s string, k int) int {
    seen := map[byte]bool{}
    count := 0
    for i := 0; i < len(s); i++ {
        seen[s[i]] = true
        if len(seen) > k {
            // start new partition
            count++
            seen = map[byte]bool{s[i]: true}
        }
    }
    if len(seen) > 0 {
        count++
    }
    return count
}

// brute force, string
// time: O(n^2)
// space: O(n)
