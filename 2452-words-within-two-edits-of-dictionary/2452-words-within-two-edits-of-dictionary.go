func twoEditWords(queries []string, dictionary []string) []string {
    ans := []string{}
    for _, query := range queries {
        for _, dict := range dictionary {
            if query == dict {
                ans = append(ans, query)
                break
            }

            count := 0
            for i := 0; i < len(query); i++ {
                if query[i] != dict[i] {
                    count++
                }
            }

            if count <= 2 {
                ans = append(ans, query)
                break
            }
        }
    }

    return ans
}

// array, brute force, string
// time: O(qdl)
// space: O(q)