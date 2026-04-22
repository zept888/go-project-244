package formatters

import (
	"code/internal/diff"
	"fmt"
	"sort"
	"strings"
)

func stylish(nodes []diff.Node) string {
	return renderNodes(nodes, 1)
}

func renderNodes(nodes []diff.Node, depth int) string {
	var lines []string
	for _, node := range nodes {
		lines = append(lines, renderNode(node, depth))
	}
	indent := strings.Repeat("    ", depth-1)
	return "{\n" + strings.Join(lines, "\n") + "\n" + indent + "}"
}

func renderNode(node diff.Node, depth int) string {
	indent := strings.Repeat("    ", depth-1)

	switch node.Type {
	case diff.Nested:
		return fmt.Sprintf("%s    %s: %s", indent, node.Key, renderNodes(node.Children, depth+1))
	case diff.Unchanged:
		return fmt.Sprintf("%s    %s: %s", indent, node.Key, formatValue(node.OldValue, depth))
	case diff.Added:
		return fmt.Sprintf("%s  + %s: %s", indent, node.Key, formatValue(node.NewValue, depth))
	case diff.Removed:
		return fmt.Sprintf("%s  - %s: %s", indent, node.Key, formatValue(node.OldValue, depth))
	case diff.Updated:
		removed := fmt.Sprintf("%s  - %s: %s", indent, node.Key, formatValue(node.OldValue, depth))
		added := fmt.Sprintf("%s  + %s: %s", indent, node.Key, formatValue(node.NewValue, depth))
		return removed + "\n" + added
	}
	return ""
}

func formatValue(val any, depth int) string {
	switch v := val.(type) {
	case map[string]any:
		return renderMap(v, depth+1)
	case nil:
		return "null"
	case bool:
		if v {
			return "true"
		}
		return "false"
	default:
		return fmt.Sprintf("%v", v)
	}
}

func renderMap(m map[string]any, depth int) string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	indent := strings.Repeat("    ", depth-1)
	var lines []string
	for _, k := range keys {
		lines = append(lines, fmt.Sprintf("%s    %s: %s", indent, k, formatValue(m[k], depth)))
	}
	return "{\n" + strings.Join(lines, "\n") + "\n" + indent + "}"
}
