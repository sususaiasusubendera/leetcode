func answerString(word string, numFriends int) string {
	if numFriends == 1 {
		return word
	}
	n := len(word)
	res := ""
	for i := 0; i < n; i++ {
		res = max(res, word[i:min(i+n-numFriends+1, n)])
	}
	return res
}

// editorial
// notice me senpai!!!