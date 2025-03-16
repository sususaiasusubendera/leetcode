func repairCars(ranks []int, cars int) int64 {
    low, high := int64(1), int64(^uint(0) >> 1)
    for low < high {
        mid := low + ((high - low) / 2)
        carsRepaired := int64(0)
        for _, r := range ranks {
            carsRepaired += int64(math.Floor(math.Sqrt(float64(mid) / float64(r))))
        }

        if carsRepaired >= int64(cars) {
            high = mid
            continue
        }
        low = mid + 1
    }

    return low
}

// NOTICE ME SENPAI!!!