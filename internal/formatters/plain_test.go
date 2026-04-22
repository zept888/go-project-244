package formatters

import (
	"code/internal/diff"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlainAdded(t *testing.T) {
	nodes := []diff.Node{
		{Key: "verbose", Type: diff.Added, NewValue: true},
	}
	result, err := Format(nodes, "plain")
	assert.NoError(t, err)
	assert.Equal(t, "Property 'verbose' was added with value: true", result)
}

func TestPlainRemoved(t *testing.T) {
	nodes := []diff.Node{
		{Key: "proxy", Type: diff.Removed, OldValue: "123.234.53.22"},
	}
	result, err := Format(nodes, "plain")
	assert.NoError(t, err)
	assert.Equal(t, "Property 'proxy' was removed", result)
}

func TestPlainUpdated(t *testing.T) {
	nodes := []diff.Node{
		{Key: "timeout", Type: diff.Updated, OldValue: 50, NewValue: 20},
	}
	result, err := Format(nodes, "plain")
	assert.NoError(t, err)
	assert.Equal(t, "Property 'timeout' was updated. From 50 to 20", result)
}

func TestPlainComplexValue(t *testing.T) {
	nodes := []diff.Node{
		{Key: "setting5", Type: diff.Added, NewValue: map[string]any{"key5": "value5"}},
	}
	result, err := Format(nodes, "plain")
	assert.NoError(t, err)
	assert.Equal(t, "Property 'setting5' was added with value: [complex value]", result)
}

func TestPlainNullValue(t *testing.T) {
	nodes := []diff.Node{
		{Key: "setting3", Type: diff.Updated, OldValue: true, NewValue: nil},
	}
	result, err := Format(nodes, "plain")
	assert.NoError(t, err)
	assert.Equal(t, "Property 'setting3' was updated. From true to null", result)
}

func TestPlainNested(t *testing.T) {
	nodes := []diff.Node{
		{Key: "common", Type: diff.Nested, Children: []diff.Node{
			{Key: "follow", Type: diff.Added, NewValue: false},
			{Key: "setting2", Type: diff.Removed, OldValue: 200},
			{Key: "setting6", Type: diff.Nested, Children: []diff.Node{
				{Key: "ops", Type: diff.Added, NewValue: "vops"},
			}},
		}},
	}
	result, err := Format(nodes, "plain")
	assert.NoError(t, err)
	assert.Equal(t, `Property 'common.follow' was added with value: false
Property 'common.setting2' was removed
Property 'common.setting6.ops' was added with value: 'vops'`, result)
}

func TestPlainUnchangedIgnored(t *testing.T) {
	nodes := []diff.Node{
		{Key: "host", Type: diff.Unchanged, OldValue: "hexlet.io"},
	}
	result, err := Format(nodes, "plain")
	assert.NoError(t, err)
	assert.Equal(t, "", result)
}
