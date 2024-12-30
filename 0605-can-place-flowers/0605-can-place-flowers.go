func canPlaceFlowers(flowerbed []int, n int) bool {
    if n == 0 {
        return true
    }

    if len(flowerbed) == 1 && flowerbed[0] == 0 {
        return true
    }

    for i := 0; i < len(flowerbed); i++ {
        if flowerbed[i] != 1 {
            if i == 0 {
                if flowerbed[i+1] == 0 {
                    n--
                    flowerbed[i] = 1
                }
            } else if i == len(flowerbed) - 1 {
                if flowerbed[i-1] == 0 {
                    n--
                    flowerbed[i] = 1
                }
            } else {
                if flowerbed[i-1] == 0 && flowerbed[i + 1] == 0 {
                    n--
                    flowerbed[i] = 1
                }
            }
        }
    }

    if n <= 0 {
        return true
    } else {
        return false
    }
}