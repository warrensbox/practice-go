func checkRecord(s string) bool {
	aCount := 0
	lCount := 0

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'A':
			lCount = 0
			aCount++
		case 'L':
			lCount++
		case 'P':
			lCount = 0
		}
		if aCount > 1 || lCount > 2 {
			return false
		}
	}

	return true
}