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

	// проверяем что результат валидный json
	var parsed []jsonNode
	assert.NoError(t, json.Unmarshal([]byte(result), &parsed))
	assert.Equal(t, 5, len(parsed))
	assert.Equal(t, "host", parsed[0].Key)
	assert.Equal(t, "unchanged", parsed[0].Type)
	assert.Equal(t, "timeout", parsed[1].Key)
	assert.Equal(t, "updated", parsed[1].Type)
	assert.Equal(t, "nested", parsed[4].Key)
	assert.Equal(t, 1, len(parsed[4].Children))
}
