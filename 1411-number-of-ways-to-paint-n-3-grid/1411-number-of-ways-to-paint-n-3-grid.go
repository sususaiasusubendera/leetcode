func numOfWays(n int) int {
    // let x y z; x != y, y != z
    // type a: x == z
    // type b: x != z
	if n == 1 {
        // 6 type a + 6 type b
		return 12
	}

    const MOD = 1_000_000_007

    a := 6
    b := 6
    for i := 2; i <= n; i++ {
        newA := (3 * a + 2 * b) % MOD
        newB := (2 * a + 2 * b) % MOD

        a, b = newA, newB
    }

    return (a + b) % MOD
}

// dp
// time: O(n)
// space: O(1)

/*
    type a
    R Y R 
    R G R
    Y R Y
    Y G Y
    G Y G
    G R G

    type b
    R Y G | G Y R
    G R Y | Y R G
    Y G R | R G Y
    
    ---

    type a (x == z)
    e.g. R Y R
    R Y R
    x y z

    x != R
    y != Y
    z != R
    x != y
    y != z

    x = {Y, G}
    y = {R, G}
    z = {G, Y}

    x = Y
    Y R G type b
    Y R Y type a
    Y G Y type a

    x = G
    G R G type a
    G R Y type b

    1a -> 3a + 2b

    ---

    type b (x != y, y != z, x != z)
    e.g. R Y G
    R Y G
    x y z

    x != R
    y != Y
    z != G
    x != y
    y != z

    x = {Y, G}
    y = {R, G}
    z = {R, Y}

    x = Y
    Y R Y type a
    Y G R type b
    Y G Y type a

    x = G
    G R Y type b

    1b -> 2a + 2b
*/