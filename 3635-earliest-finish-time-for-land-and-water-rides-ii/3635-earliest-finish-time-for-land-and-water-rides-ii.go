func earliestFinishTime(landStartTime []int, landDuration []int, waterStartTime []int, waterDuration []int) int {
	landRidesEff := []LandRide{}
	landRidesSortedStart := []LandRide{}
	for i := 0; i < len(landStartTime); i++ {
		landRidesEff = append(landRidesEff, LandRide{landStartTime[i], landDuration[i], landStartTime[i] + landDuration[i]})
		landRidesSortedStart = append(landRidesSortedStart, LandRide{landStartTime[i], landDuration[i], landStartTime[i] + landDuration[i]})
	}

	waterRidesEff := []WaterRide{}
	waterRidesSortedStart := []WaterRide{}
	for i := 0; i < len(waterStartTime); i++ {
		waterRidesEff = append(waterRidesEff, WaterRide{waterStartTime[i], waterDuration[i], waterStartTime[i] + waterDuration[i]})
		waterRidesSortedStart = append(waterRidesSortedStart, WaterRide{waterStartTime[i], waterDuration[i], waterStartTime[i] + waterDuration[i]})
	}

	sort.Slice(landRidesEff, func(i, j int) bool { return landRidesEff[i].end < landRidesEff[j].end })
	sort.Slice(waterRidesEff, func(i, j int) bool { return waterRidesEff[i].end < waterRidesEff[j].end })

	sort.Slice(landRidesSortedStart, func(i, j int) bool { return landRidesSortedStart[i].start < landRidesSortedStart[j].start })
	sort.Slice(waterRidesSortedStart, func(i, j int) bool { return waterRidesSortedStart[i].duration < waterRidesSortedStart[j].start })


    // LAND -> WATER
    landRes := 1_000_000_007
    for i := 0; i < len(waterRidesSortedStart); i++ {
        if waterRidesSortedStart[i].start < landRidesEff[0].end {
            landRes = min(landRes, landRidesEff[0].end + waterRidesSortedStart[i].duration)
        } else {
            landRes = min(landRes, waterRidesSortedStart[i].start + waterRidesSortedStart[i].duration)
        }
    }

    // WATER TO LAND
    waterRes := 1_000_000_007
    for i := 0; i < len(landRidesSortedStart); i++ {
        if landRidesSortedStart[i].start < waterRidesEff[0].end {
            landRes = min(landRes, waterRidesEff[0].end + landRidesSortedStart[i].duration)
        } else {
            landRes = min(landRes, landRidesSortedStart[i].start + landRidesSortedStart[i].duration)
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