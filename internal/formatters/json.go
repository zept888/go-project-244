package formatters

import (
	"code/internal/diff"
	"encoding/json"
	"fmt"
)

type jsonNode struct {
	Key      string     `json:"key"`
	Type     string     `json:"type"`
	OldValue any        `json:"oldValue,omitempty"`
	NewValue any        `json:"newValue,omitempty"`
	Children []jsonNode `json:"children,omitempty"`
}

func toJSONNodes(nodes []diff.Node) []jsonNode {
	result := make([]jsonNode, 0, len(nodes))
	for _, n := range nodes {
		jn := jsonNode{
			Key:      n.Key,
			Type:     string(n.Type),
			OldValue: n.OldValue,
			NewValue: n.NewValue,
		}
		if len(n.Children) > 0 {
			jn.Children = toJSONNodes(n.Children)
		}
		result = append(result, jn)
	}
	return result
}

func formatJSON(nodes []diff.Node) (string, error) {
	data, err := json.MarshalIndent(toJSONNodes(nodes), "", "    ")
	if err != nil {
		return "", fmt.Errorf("json marshal error: %w", err)
	}
	return string(data), nil
}
