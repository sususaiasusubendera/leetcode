func validateCoupons(code []string, businessLine []string, isActive []bool) []string {
	mBusiness := map[string]bool{
		"electronics": true,
		"grocery":     true,
		"pharmacy":    true,
		"restaurant":  true,
	}

	temp := []Record{}
	for i := range code {
		validCode, validBusiness := true, true

		// code
		if len(code[i]) == 0 {
			validCode = false
		} else {
			for _, c := range code[i] {
				if ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') || ('0' <= c && c <= '9') || c == '_' {
					continue
				} else {
					validCode = false
					break
				}
			}
		}

		// business
		if !mBusiness[businessLine[i]] {
			validBusiness = false
		}

		if validCode && validBusiness && isActive[i] {
			temp = append(temp, Record{code[i], businessLine[i][0]})
		}
	}

	sort.Slice(temp, func(i, j int) bool {
		if temp[i].business == temp[j].business {
			return temp[i].code < temp[j].code
		}

		return temp[i].business < temp[j].business
	})

	ans := make([]string, len(temp))
	for i := range temp {
		ans[i] = temp[i].code
	}

	return ans
}

type Record struct {
	code     string
	business byte
}

// array, hash map, sorting, string
// time: O(nm + nlog(n))
// space: O(n)