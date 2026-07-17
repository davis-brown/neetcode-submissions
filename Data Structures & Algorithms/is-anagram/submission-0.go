func isAnagram(s string, t string) bool {
	if(len(s) != len(t)) {
		return false
	}
	
	tally := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		tally[s[i]]++
	}

	for j := 0; j < len(t); j++ {
		tally[t[j]]--
		if(tally[t[j]] < 0) {
			return false
		}

	}

	return true
}
