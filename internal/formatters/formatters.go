package formatters

import (
	"code/internal/diff"
	"fmt"
)

func Format(nodes []diff.Node, format string) (string, error) {
	switch format {
	case "stylish", "":
		return stylish(nodes), nil
	case "plain":
		return plain(nodes), nil
	default:
		return "", fmt.Errorf("unknown format: %q", format)
	}
}
