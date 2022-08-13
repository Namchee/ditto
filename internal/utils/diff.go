package utils

import (
	"github.com/Namchee/ditto/internal/entity"
	"github.com/r3labs/diff/v3"
)

// HasDiff checks if a runner result has different value than first element
func HasDiff(
	resp []*entity.FetchResult,
	config *entity.Configuration,
) (bool, []diff.Changelog) {
	changelogs := []diff.Changelog{}

	for idx := 1; idx < len(resp); idx++ {
		changes, _ := diff.Diff(resp[0].Response, resp[idx].Response)

		changelogs = append(changelogs, changes)
	}

	hasDiff := false
	for _, d := range changelogs {
		hasDiff = len(d) > 0
	}

	return hasDiff, changelogs
}
