func angleClock(hour int, minutes int) float64 {
    minutesDeg := (float64(minutes) / 5) * 30
    hourDeg := math.Mod((float64(hour) + (float64(minutes) / 60)) * 30, 360)
    deg := abs(minutesDeg - hourDeg)
    if deg >= 180 {
        return 360 - deg
    }
    return deg
}

func abs(n float64) float64 {
    if n < 0 {
        return -n
    }
    return n
}

// math
// time: O(1)
// space: O(1)