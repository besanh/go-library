package util

import (
	"encoding/json"
	"fmt"
	"strings"
)

// DecodeUnicode decodes Unicode escape sequences in a string.
// For example, it converts `\u003c` to `<`.
// Returns the decoded string or an error if decoding fails.
func (i *Util) DecodeUnicode(text string) (string, error) {
	jsonStr := fmt.Sprintf(`"%s"`, text)
	var output string
	if err := json.Unmarshal([]byte(jsonStr), &output); err != nil {
		return "", err
	}
	return output, nil
}

// JoinWithSeparator joins elements of a slice with a separator.
func (i *Util) JoinWithSeparator(input []string, separator ...string) string {
	return strings.Join(input, getSeparator(separator))
}

// SplitWithSeparator splits a string into a slice using a separator.
func (i *Util) SplitWithSeparator(input string, separator ...string) []string {
	return strings.Split(input, getSeparator(separator))
}

// getSeparator returns the first separator if provided or default to "\n".
func getSeparator(separator []string) string {
	if len(separator) > 0 && separator[0] != "" {
		return separator[0]
	}
	return "\n"
}
