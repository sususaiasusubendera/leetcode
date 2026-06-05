func earliestFinishTime(landStartTime []int, landDuration []int, waterStartTime []int, waterDuration []int) int {
	landRides := []LandRide{}
	for i := 0; i < len(landStartTime); i++ {
		landRides = append(landRides, LandRide{landStartTime[i], landDuration[i], landStartTime[i] + landDuration[i]})
	}

	waterRides := []WaterRide{}
	for i := 0; i < len(waterStartTime); i++ {
		waterRides = append(waterRides, WaterRide{waterStartTime[i], waterDuration[i], waterStartTime[i] + waterDuration[i]})
	}

	sort.Slice(landRides, func(i, j int) bool { return landRides[i].end < landRides[j].end })
	sort.Slice(waterRides, func(i, j int) bool { return waterRides[i].end < waterRides[j].end })

    // LAND -> WATER
    landRes := 1_000_000_007
    for i := 0; i < len(waterRides); i++ {
        if waterRides[i].start < landRides[0].end {
            landRes = min(landRes, landRides[0].end + waterRides[i].duration)
        } else {
            landRes = min(landRes, waterRides[i].start + waterRides[i].duration)
        }
    }

    // WATER TO LAND
    waterRes := 1_000_000_007
    for i := 0; i < len(landRides); i++ {
        if landRides[i].start < waterRides[0].end {
            landRes = min(landRes, waterRides[0].end + landRides[i].duration)
        } else {
            landRes = min(landRes, landRides[i].start + landRides[i].duration)
        }
    }

    return min(landRes, waterRes)
}

type LandRide struct {
	start    int
	duration int
	end      int
}

type WaterRide struct {
	start    int
	duration int
	end      int
}

// array, greedy, sorting
// time: O(nlog(n) + mlog(m) + n + m)
// space: O(n + m)

// note
// actually, sorting is unnecessary