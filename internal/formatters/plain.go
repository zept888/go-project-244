package formatters

import (
	"code/internal/diff"
	"fmt"
	"strings"
)

func plain(nodes []diff.Node) string {
	return renderPlain(nodes, "")
}

func renderPlain(nodes []diff.Node, path string) string {
	var lines []string
	for _, node := range nodes {
		currentPath := buildPath(path, node.Key)
		switch node.Type {
		case diff.Nested:
			lines = append(lines, renderPlain(node.Children, currentPath))
		case diff.Added:
			lines = append(lines, fmt.Sprintf("Property '%s' was added with value: %s", currentPath, formatPlainValue(node.NewValue)))
		case diff.Removed:
			lines = append(lines, fmt.Sprintf("Property '%s' was removed", currentPath))
		case diff.Updated:
			lines = append(lines, fmt.Sprintf("Property '%s' was updated. From %s to %s", currentPath, formatPlainValue(node.OldValue), formatPlainValue(node.NewValue)))
		}
	}
	return strings.Join(lines, "\n")
}

func buildPath(parent, key string) string {
	if parent == "" {
		return key
	}
	return parent + "." + key
}

func formatPlainValue(val any) string {
	switch v := val.(type) {
	case map[string]any:
		return "[complex value]"
	case nil:
		return "null"
	case bool:
		if v {
			return "true"
		}
		return "false"
	case string:
		return fmt.Sprintf("'%s'", v)
	default:
		return fmt.Sprintf("%v", v)
	}
}
