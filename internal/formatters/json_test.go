package formatters

import (
	"code/internal/diff"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatJSON(t *testing.T) {
	nodes := []diff.Node{
		{Key: "host", Type: diff.Unchanged, OldValue: "hexlet.io"},
		{Key: "timeout", Type: diff.Updated, OldValue: 50, NewValue: 20},
		{Key: "proxy", Type: diff.Removed, OldValue: "123.234.53.22"},
		{Key: "verbose", Type: diff.Added, NewValue: true},
		{Key: "nested", Type: diff.Nested, Children: []diff.Node{
			{Key: "key", Type: diff.Unchanged, OldValue: "value"},
		}},
	}

	result, err := Format(nodes, "json")
	assert.NoError(t, err)

	var parsed map[string]any
	assert.NoError(t, json.Unmarshal([]byte(result), &parsed))
	assert.Contains(t, parsed, "diff")
	assert.NotEmpty(t, parsed["diff"])
}
