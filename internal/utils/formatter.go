package utils

import (
	"fmt"

	"github.com/Namchee/ditto/internal/constant"
	"github.com/Namchee/ditto/internal/entity"
)

// FormatResult apply formatting to test result
func FormatResult(result *entity.RunnerResult, status bool) string {
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

	if !status {
		emoji = constant.FailEmoji
		text = constant.FailText
		format += ". Please check the generated test log."
	}

	return fmt.Sprintf(format, result.Name, emoji, text)
}
