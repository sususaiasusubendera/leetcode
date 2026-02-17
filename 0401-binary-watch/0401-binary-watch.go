func readBinaryWatch(turnedOn int) []string {
	ans := []string{}

	var dfs func(pos, countOn, h, m int)
	dfs = func(pos, countOn, h, m int) {
        if h > 11 || m > 59 { return }

        if pos == 10 {
            if countOn == turnedOn {
                ans = append(ans, fmt.Sprintf("%d:%02d", h, m))
            }
            return
        }

        // LED off
        dfs(pos+1, countOn, h, m)

        // LED on
        if pos < 4 {
            dfs(pos+1, countOn+1, h+(1<<pos), m)
        } else {
            dfs(pos+1, countOn+1, h, m+(1<<(pos-4)))
        }
	}

    dfs(0, 0, 0, 0)
    return ans
}

// backtracking
// time: O(1)
// space: O(1)