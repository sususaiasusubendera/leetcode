func constructDistancedSequence(n int) []int {
    size := (n-1) * 2 + 1
    inserted := make([]bool, n+1)
    var backtrack func(cand []int, idx int) []int
    backtrack = func(cand []int, idx int) []int {
        if idx >= size {
            return cand
        }
        if cand[idx] != 0 {
            return backtrack(cand, idx+1)
        }

        for i := n; i >= 1; i-- {
            if inserted[i] { 
                continue
            }
            if i != 1 && (idx+i >= size || cand[idx+i] != 0) { 
                continue
            }
            
            inserted[i] = true
            cand[idx] = i
            if i != 1 {
                cand[idx+i] = i
            }
            if res := backtrack(cand,idx+1); len(res) > 0 {
                return res
            }
            cand[idx]= 0
            if i != 1 {
                cand[idx+i] = 0
            }
            inserted[i] = false

        }

        return nil
    }

    return backtrack(make([]int,size), 0)
}

// notice me senpai