package formatters

import (
	"code/internal/diff"
	"fmt"
)

// Format formats the diff nodes according to the specified format.
// Supported formats: stylish, plain, json. Defaults to stylish if empty.
func Format(nodes []diff.Node, format string) (string, error) {
	switch format {
	case "stylish", "":
		return stylish(nodes), nil
	case "plain":
		return plain(nodes), nil
	case "json":
		return formatJSON(nodes)
	default:
		return "", fmt.Errorf("unknown format: %q", format)
	}
}
