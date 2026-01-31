func nextGreatestLetter(letters []byte, target byte) byte {
	low, high := 0, len(letters)
	for low < high {
		mid := low + ((high - low) / 2)
		if letters[mid] > target {
			high = mid
		} else {
			low = mid + 1
		}
	}

	if low < len(letters) {
		return letters[low]
	} else {
		return letters[0]
	}
}

// array, binary search
// time: O(log(n))
// space: O(1)