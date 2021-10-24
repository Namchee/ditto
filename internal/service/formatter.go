package service

import (
	"fmt"

	"github.com/Namchee/ditto/internal/constant"
	"github.com/Namchee/ditto/internal/utils"
)

// FormatResult apply formatting to test result
func FormatResult(test string, diff []int) string {
	emoji := constant.PassEmoji
	text := constant.PassText
	format := "%s: %s %s"

	if len(diff) != 0 {
		emoji = constant.FailEmoji
		text = constant.FailText
		format += fmt.Sprintf(
			" = Endpoint(s) with index %s have different result(s)",
			utils.IntSliceToString(diff),
		)
	}

	return fmt.Sprintf(format, test, emoji, text)
}
