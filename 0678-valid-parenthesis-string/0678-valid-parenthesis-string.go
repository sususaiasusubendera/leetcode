func checkValidString(s string) bool {
    minBalance, maxBalance := 0, 0
    for _, c := range s {
        if c == '(' {
            minBalance++
            maxBalance++
        } else if c == ')' {
            minBalance--
            maxBalance--
        } else { // c == '*'
            minBalance--
            maxBalance++
        }

        // not valid 
        if maxBalance < 0 {
            return false
        }

        // clip negative-balance for minBalance
        minBalance = max(minBalance, 0)
    }

    return minBalance == 0 // valid paranthesis when balance is 0
}

// greedy, string
// a beautiful and clean approach form la_castille
// time: O(n)
// space: O(1)