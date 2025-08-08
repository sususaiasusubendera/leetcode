func soupServings(n int) float64 {
    if n > 5000 {
        return 1
    }
    cache := make(map[int]float64)
    probability := calculateProbability(n, n, cache)
    return probability
}

// notice me senpai!!!
// dp

func calculateProbability(a int, b int, cache map[int]float64) float64 {
    key := getKey(a, b)
    value, exists := cache[key]
    if exists {
        return value
    }
    if a <= 0 && b <= 0 {
        return 0.5
    }
    if a <= 0 && b > 0 {
        return 1
    }
    if a > 0 && b <= 0 {
        return 0
    }
    p1 := calculateProbability(a - 100, b, cache)
    p2 := calculateProbability(a - 75, b - 25, cache)
    p3 := calculateProbability(a - 50, b - 50, cache)
    p4 := calculateProbability(a - 25, b - 75, cache)
    result := 0.25 * (p1 + p2 + p3 + p4)
    cache[key] = result
    return result
}

func getKey(a int, b int) int {
    return (a << 15) + b
}