package utils

import (
	"fmt"
	"strings"
)

// IntSliceToString returns a string representation of a integer slice
func IntSliceToString(slice []int) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(slice)), ", "), "[]")
}
