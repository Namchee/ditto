package service

// GetDiff returns list of item index that have different value than first element
func GetDiff(resp []string) []int {
	var diffs []int

	for i := 1; i < len(resp); i++ {
		if resp[i] != resp[0] {
			diffs = append(diffs, i)
		}
	}
	return diffs
}
