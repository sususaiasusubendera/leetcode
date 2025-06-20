func maxDistance(s string, k int) int {
	latitude, longitude, ans := 0, 0, 0
	n := len(s)
	for i := 0; i < n; i++ {
		switch s[i] {
		case 'N':
			latitude++
		case 'S':
			latitude--
		case 'E':
			longitude++
		case 'W':
			longitude--
		}
		ans = max(ans, min(abs(latitude)+abs(longitude)+k*2, i+1))
	}
	return ans
}

// NOTICE ME SENPAI!!!

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}