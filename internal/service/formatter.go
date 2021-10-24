package service

import (
	"fmt"

	"github.com/Namchee/ditto/internal/constant"
	"github.com/Namchee/ditto/internal/entity"
	"github.com/Namchee/ditto/internal/utils"
)

// FormatResult apply formatting to test result
func FormatResult(result *entity.TestResult) string {
	if result.Error != nil {
		return fmt.Sprintf(
			"%s: %s %s = Failed to run test: %s",
			result.Name,
			constant.FailEmoji,
			constant.FailText,
			result.Error.Error(),
		)
	}

	emoji := constant.PassEmoji
	text := constant.PassText
	format := "%s: %s %s"

	if len(result.Diff) != 0 {
		emoji = constant.FailEmoji
		text = constant.FailText
		format += fmt.Sprintf(
			" = Endpoint(s) with index %s have different result(s)",
			utils.IntSliceToString(result.Diff),
		)
	}

	return fmt.Sprintf(format, result.Name, emoji, text)
}
