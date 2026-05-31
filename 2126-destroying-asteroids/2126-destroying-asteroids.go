func asteroidsDestroyed(mass int, asteroids []int) bool {
    sort.Ints(asteroids)
    for i := 0; i < len(asteroids); i++ {
        if mass >= asteroids[i] {
            mass += asteroids[i]
        } else {
            return false
        }
    }

    return true
}

// array, sorting
// time: O(nlog(n))
// space: O(1)