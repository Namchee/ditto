package utils

import (
	"fmt"

	"github.com/Namchee/ditto/internal/constant"
	"github.com/Namchee/ditto/internal/entity"
)

// FormatResult apply formatting to test result
func FormatResult(result *entity.RunnerResult, status bool) string {
	emoji := constant.PassEmoji
	text := constant.PassText
	format := "%s: %s %s"

	if result.Error != nil || !status {
		emoji = constant.FailEmoji
		text = constant.FailText
		format += ". Please check the generated test log."
	}

	return fmt.Sprintf(format, result.Name, emoji, text)
}
