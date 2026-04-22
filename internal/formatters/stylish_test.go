package formatters

import (
	"code/internal/diff"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatUnknown(t *testing.T) {
	_, err := Format([]diff.Node{}, "unknown")
	assert.Error(t, err)
}

func TestFormatStylish(t *testing.T) {
	nodes := []diff.Node{
		{Key: "follow", Type: diff.Removed, OldValue: false},
		{Key: "host", Type: diff.Unchanged, OldValue: "hexlet.io"},
		{Key: "timeout", Type: diff.Updated, OldValue: 50, NewValue: 20},
		{Key: "proxy", Type: diff.Removed, OldValue: "123.234.53.22"},
		{Key: "verbose", Type: diff.Added, NewValue: true},
		{Key: "nested", Type: diff.Nested, Children: []diff.Node{
			{Key: "key", Type: diff.Unchanged, OldValue: "value"},
		}},
	}

	result, err := Format(nodes, "stylish")
	assert.NoError(t, err)
	assert.Equal(t, `{
  - follow: false
    host: hexlet.io
  - timeout: 50
  + timeout: 20
  - proxy: 123.234.53.22
  + verbose: true
    nested: {
        key: value
    }
}`, result)
}
