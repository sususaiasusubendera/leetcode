func numberOfWays(corridor string) int {
    const MOD = 1_000_000_007

    seat := 0 // count seat
    plant := 0 // count plant
    dp := 1
    for i := range corridor {
        if corridor[i] == 'S' {
            seat++

            if seat == 2 {
                plant = 0
            }
        } else { // P
            if seat == 2 {
                plant++
            }
        }

        if seat == 3 {
            dp = dp * (plant + 1) % MOD
            seat = 1
            plant = 0
        }
    }

    if seat != 2 {
        return 0
    }

    return dp
}

// p.s. the greedy approach is derived from the dp approach
// greedy, string, dp (bot-up rolling state)
// time: O(n)
// space: O(1)