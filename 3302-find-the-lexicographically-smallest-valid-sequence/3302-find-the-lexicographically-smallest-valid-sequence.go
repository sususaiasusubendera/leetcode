func validSequence(word1 string, word2 string) []int {
	n, m := len(word1), len(word2)

	// store index of word1 that has same char with word2, lexicographically smallest seq
	last := make([]int, m)
	for i := range last {
		last[i] = -1
	}

	j := m - 1
    for i := n - 1; i >= 0; i-- {
        if j >= 0 && word1[i] == word2[j] {
            last[j] = i
            j -= 1
        }
    }

    ans := []int{}
    skip := 0
    j = 0
    for i := 0; i < n; i++ {
        if j == m {
            break
        }
        
        // here's the thing: greedily use the skip (mismatch) as soon as the conditions are met
        // i doubted it at first, but when i think about it again visually in my brain, it make sense
        // use skip when: 
        // 1) you still have a skip (ofc, cuh!)
        // 2) j is the last index of word2 OR the index i of word1 is valid (i < last[j+i])
        if word1[i] == word2[j] || (skip == 0 && (j == m-1 || i < last[j+1])) {
            // use skip
            if word1[i] != word2[j] {
                skip += 1
            }
            ans = append(ans, i)
            j += 1
        }
    }

    if j == m {
        return ans
    }

    return []int{}
}

// greedy (very tricky), string, two pointers
// time: O(n + m)
// space: O(m)