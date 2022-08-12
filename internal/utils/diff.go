package utils

import (
	"github.com/Namchee/ditto/internal/entity"
	"github.com/r3labs/diff/v3"
)

// HasDiff checks if a runner result has different value than first element
func HasDiff(
	resp []*entity.FetchResult,
	config *entity.Configuration,
) []diff.Changelog {
	changes, err := diff.Diff(resp[0], resp[1])

	changes.
}
