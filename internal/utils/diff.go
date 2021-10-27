package utils

import "github.com/Namchee/ditto/internal/entity"

// HasDiff checks if a runner result has different value than first element
func HasDiff(resp []*entity.FetchResult, config *entity.Configuration) bool {
	for i := 1; i < len(resp); i++ {
		if resp[i].Response != resp[0].Response ||
			(config.Status && resp[i].Status != resp[0].Status) {
			return true
		}
	}

	return false
}
