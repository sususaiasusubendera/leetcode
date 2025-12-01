func maxRunTime(n int, batteries []int) int64 {
    sum := 0
    for _, battery := range batteries {
        sum += battery
    }

    // idea: 
    // check if all n computers can run for T (target) minutes
    // each battery contributes min(battery[i], T); let total contribution of batteries = totBat
    // T is feasible if totBat / n >= T
    low, high := 0, sum / n
    for low < high {
        target := (low + high + 1) / 2 // upper mid

        extra := 0
        for _, battery := range batteries {
            if battery >= target {
                extra += target
            } else {
                extra += battery
            }
        }

        if extra / n >= target {
            low = target
        } else {
            high = target - 1
        }
    }

    return int64(low)
}

// array, binary search
// time: O(mlog(sum/n))
// space: O(1)